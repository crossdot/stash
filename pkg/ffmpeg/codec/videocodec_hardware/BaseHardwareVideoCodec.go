package videocodec_hardware

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
	"github.com/stashapp/stash/pkg/logger"
	"github.com/stashapp/stash/pkg/models"
)

type BaseHardwareVideoCodec struct {
	self codec.VideoCodec
}

func (b *BaseHardwareVideoCodec) InitHWSupport(ctx context.Context) bool {
	if b.self == nil {
		panic("hardware codec not initialized correctly")
	}

	const hwInitLogTimeoutSecondsDefault = 5
	hwInitLogTimeout := time.Duration(hwInitLogTimeoutSecondsDefault) * time.Second

	done := make(chan bool)

	// run initialization in a separate goroutine
	go func() {
		result := b.initHWSupport(ctx)
		done <- result
		close(done)
	}()

	timer := time.NewTimer(hwInitLogTimeout)
	defer timer.Stop()

	select {
	case success := <-done:
		// initialization finished before timeout
		return success
	case <-timer.C:
		// initialization timed out
		logger.Warnf("[InitHWSupport] Hardware codec initialization is taking longer than %s...", hwInitLogTimeout)
		logger.Info("[InitHWSupport] Hardware encoding will not be available until initialization is complete.")
		return false
	}
}

func (b *BaseHardwareVideoCodec) initHWSupport(ctx context.Context) bool {
	const minHeight int = 480

	var args ffmpeg.Args
	args = append(args, "-hide_banner")
	// args = args.LogLevel(LogLevelWarning)
	args = b.self.(codec.HardwareVideoCodec).HwDeviceInit(args, false)
	args = args.Format("lavfi")
	vFile := &models.VideoFile{Width: 1280, Height: 720}
	args = args.Input(fmt.Sprintf("color=c=red:s=%dx%d", vFile.Width, vFile.Height))
	args = args.Duration(0.1)

	// Test scaling
	videoFilter := b.hwMaxResFilter(vFile, minHeight, false)
	args = append(args, b.self.CodecInit()...)
	args = args.VideoFilter(videoFilter)

	args = args.Format("null")
	args = args.Output("-")

	// #6064 - add timeout to context to prevent hangs
	const hwTestTimeoutSecondsDefault = 10
	hwTestTimeoutSeconds := hwTestTimeoutSecondsDefault * time.Second
	// allow timeout to be overridden with environment variable
	if timeout := os.Getenv("STASH_HW_TEST_TIMEOUT"); timeout != "" {
		if seconds, err := strconv.Atoi(timeout); err == nil {
			hwTestTimeoutSeconds = time.Duration(seconds) * time.Second
		}
	}

	testCtx, cancel := context.WithTimeout(ctx, hwTestTimeoutSeconds)
	defer cancel()

	cmd := b.Command(testCtx, args)
	cmd.WaitDelay = time.Second
	logger.Tracef("[InitHWSupport] Testing codec %s: %v", b.self, cmd.Args)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if testCtx.Err() != nil {
			logger.Debugf("[InitHWSupport] Codec %s test timed out after %s", b.self, hwTestTimeoutSeconds)
			return false
		}

		errOutput := stderr.String()
		if len(errOutput) == 0 {
			errOutput = err.Error()
		}
		logger.Debugf("[InitHWSupport] Codec %s not supported. Error output:\n%s", b.self, errOutput)
		return false
	} else {
		return true
	}
}

func (b *BaseHardwareVideoCodec) hwCanFullHWTranscode(ctx context.Context, vf *models.VideoFile, reqHeight int) bool {
	var args ffmpeg.Args
	args = append(args, "-hide_banner")
	args = args.LogLevel(ffmpeg.LogLevelWarning)
	args = args.XError()
	args = b.self.(codec.HardwareVideoCodec).HwDeviceInit(args, true)
	args = args.Input(vf.Path)
	args = args.Duration(1)

	videoFilter := b.hwMaxResFilter(vf, reqHeight, true)
	args = append(args, b.self.CodecInit()...)
	args = args.VideoFilter(videoFilter)

	args = args.Format("null")
	args = args.Output("-")

	cmd := b.Command(ctx, args)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		errOutput := stderr.String()

		if len(errOutput) == 0 {
			errOutput = err.Error()
		}

		logger.Debugf("[InitHWSupport] Full hardware transcode for file %s not supported. Error output:\n%s", vf.Basename, errOutput)
		return false
	}

	return true
}

var scaler_re = regexp.MustCompile(`scale=(?P<value>([-\d]+):([-\d]+))`)

func templateReplaceScale(input string, template string, match []int, vf *models.VideoFile, minusonehack bool) string {
	result := []byte{}

	if minusonehack {
		// Parse width and height
		w, err := strconv.Atoi(input[match[4]:match[5]])
		if err != nil {
			logger.Error("failed to parse width")
			return input
		}
		h, err := strconv.Atoi(input[match[6]:match[7]])
		if err != nil {
			logger.Error("failed to parse height")
			return input
		}

		// Calculate ratio
		ratio := float64(vf.Width) / float64(vf.Height)
		if w < 0 {
			w = int(math.Round(float64(h) * ratio))
		} else if h < 0 {
			h = int(math.Round(float64(w) / ratio))
		}

		// Fix not divisible by 2 errors
		if w%2 != 0 {
			w++
		}
		if h%2 != 0 {
			h++
		}

		template = strings.ReplaceAll(template, "$value", fmt.Sprintf("%d:%d", w, h))
	}

	res := string(scaler_re.ExpandString(result, template, input, match))

	matchStart := match[0]
	matchEnd := match[1]

	return input[0:matchStart] + res + input[matchEnd:]
}

// Replace video filter scaling with hardware scaling for full hardware transcoding (also fixes the format)
func (b *BaseHardwareVideoCodec) hwCodecFilter(args ffmpeg.VideoFilter, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter {
	sargs := string(args)

	match := scaler_re.FindStringSubmatchIndex(sargs)
	if match == nil {
		return b.self.(codec.HardwareVideoCodec).HwApplyFullHWFilter(args, fullhw)
	}

	return b.self.(codec.HardwareVideoCodec).HwApplyScaleTemplate(sargs, match, vf, fullhw)
}

// Return a maxres filter
func (b *BaseHardwareVideoCodec) hwMaxResFilter(vf *models.VideoFile, reqHeight int, fullhw bool) ffmpeg.VideoFilter {
	if vf.Width == 0 || vf.Height == 0 {
		return ""
	}
	videoFilter := b.self.(codec.HardwareVideoCodec).HwFilterInit(fullhw)
	maxWidth, maxHeight := b.self.(codec.HardwareVideoCodec).HwCodecMaxRes()
	videoFilter = videoFilter.ScaleMaxLM(vf.Width, vf.Height, reqHeight, maxWidth, maxHeight)
	return b.hwCodecFilter(videoFilter, vf, fullhw)
}
