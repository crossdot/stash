package videocodec_hardware

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/logger"
	"github.com/stashapp/stash/pkg/models"
)

type BaseHardwareVideoCodec struct {
	self codec.Codec
}

var scaler_re = regexp.MustCompile(`scale=(?P<value>([-\d]+):([-\d]+))`)

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
	args = b.self.(codec.HardwareCodec).HwDeviceInit(args, false)
	args = args.Format("lavfi")
	vFile := &models.VideoFile{Width: 1280, Height: 720}
	args = args.Input(fmt.Sprintf("color=c=red:s=%dx%d", vFile.Width, vFile.Height))
	args = args.Duration(0.1)

	// Test scaling
	videoFilter := b.hwMaxResFilter(vFile, minHeight, false)
	args = append(args, CodecInit(codec)...)
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
	var args Args
	args = append(args, "-hide_banner")
	args = args.LogLevel(LogLevelWarning)
	args = args.XError()
	args = b.self.(codec.HardwareCodec).HwDeviceInit(args, true)
	args = args.Input(vf.Path)
	args = args.Duration(1)

	videoFilter := b.hwMaxResFilter(vf, reqHeight, true)
	args = append(args, CodecInit(codec)...)
	args = args.VideoFilter(videoFilter)

	args = args.Format("null")
	args = args.Output("-")

	cmd := f.Command(ctx, args)

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

// func (f *BaseHardwareVideoCodec) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
// 	return args
// }

// Return a maxres filter
func (b *BaseHardwareVideoCodec) hwMaxResFilter(vf *models.VideoFile, reqHeight int, fullhw bool) VideoFilter {
	if vf.Width == 0 || vf.Height == 0 {
		return ""
	}
	videoFilter := b.self.(codec.HardwareCodec).HwFilterInit(fullhw)
	maxWidth, maxHeight := b.self.(codec.HardwareCodec).HwCodecMaxRes()
	videoFilter = videoFilter.ScaleMaxLM(vf.Width, vf.Height, reqHeight, maxWidth, maxHeight)
	return b.hwCodecFilter(videoFilter, vf, fullhw)
}

// Replace video filter scaling with hardware scaling for full hardware transcoding (also fixes the format)
func (b *BaseHardwareVideoCodec) hwCodecFilter(args VideoFilter, vf *models.VideoFile, fullhw bool) VideoFilter {
	sargs := string(args)

	match := scaler_re.FindStringSubmatchIndex(sargs)
	if match == nil {
		return b.hwApplyFullHWFilter(args, fullhw)
	}

	return b.hwApplyScaleTemplate(sargs, match, vf, fullhw)
}
