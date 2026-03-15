package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
)

type Codec_hardware_ivp9 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_ivp9() *Codec_hardware_ivp9 {
	c := &Codec_hardware_ivp9{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (c *Codec_hardware_ivp9) Name() string {
	return "VP9 Intel Quick Sync Video (QSV)"
}

func (c *Codec_hardware_ivp9) CodeName() string {
	return "vp9_qsv"
}

func (c *Codec_hardware_ivp9) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	args = append(args,
		"-global_quality", "20",
		"-preset", "faster",
	)
	return args
}

func (f *Codec_hardware_ivp9) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	if fullhw {
		args = append(args, "-hwaccel")
		args = append(args, "qsv")
		args = append(args, "-hwaccel_output_format")
		args = append(args, "qsv")
	} else {
		args = append(args, "-init_hw_device")
		args = append(args, "qsv=hw")
		args = append(args, "-filter_hw_device")
		args = append(args, "hw")
	}
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_ivp9) HwFilterInit(fullhw bool) ffmpeg.VideoFilter {
	var videoFilter ffmpeg.VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("hwupload=extra_hw_frames=64")
		videoFilter = videoFilter.Append("format=qsv")
	}
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_ivp9) HwApplyFullHWFilter(args ffmpeg.VideoFilter, fullhw bool) ffmpeg.VideoFilter {
	if fullhw && f.version.Gteq(Version{major: 3, minor: 3}) { // Added in FFMpeg 3.3
		args = args.Append("scale_qsv=format=nv12")
	}
	return args
}

// Switch scaler
func (c *Codec_hardware_ivp9) HwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter {
	var template string

	template = "scale_qsv=$value"
	if fullhw && f.version.Gteq(Version{major: 3, minor: 3}) { // Added in FFMpeg 3.3
		template += ":format=nv12"
	}

	// BUG: [scale_qsv]: Size values less than -1 are not acceptable.
	return ffmpeg.VideoFilter(templateReplaceScale(sargs, template, match, vf, true))
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_ivp9) HwCodecMaxRes() (int, int) {
	return 0, 0
}

// Return if a hardware accelerated for HLS is available
func (c *Codec_hardware_ivp9) HwCodecHLSCompatible() bool {
	return false
}

// Return if a hardware accelerated codec for MP4 is available
func (c *Codec_hardware_ivp9) HwCodecMP4Compatible() bool {
	return false
}

// Return if a hardware accelerated codec for WebM is available
func (c *Codec_hardware_ivp9) HwCodecWEBMCompatible() bool {
	return true
}

var _ codec.Codec = (*Codec_hardware_ivp9)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_ivp9)(nil)
