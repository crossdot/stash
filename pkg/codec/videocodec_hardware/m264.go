package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
)

type Codec_hardware_m264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_m264() *Codec_hardware_m264 {
	c := &Codec_hardware_m264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (c *Codec_hardware_m264) Name() string {
	return "H264 VideoToolbox"
}

func (c *Codec_hardware_m264) CodeName() string {
	return "h264_videotoolbox"
}

func (c *Codec_hardware_m264) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	args = append(args,
		"-realtime", "1",
	)
	return args
}

func (f *Codec_hardware_m264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	if fullhw {
		args = append(args, "-hwaccel")
		args = append(args, "videotoolbox")
		args = append(args, "-hwaccel_output_format")
		args = append(args, "videotoolbox_vld")
	} else {
		args = append(args, "-init_hw_device")
		args = append(args, "videotoolbox=vt")
	}
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_m264) HwFilterInit(fullhw bool) ffmpeg.VideoFilter {
	var videoFilter ffmpeg.VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload")
	}
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_m264) HwApplyFullHWFilter(args ffmpeg.VideoFilter, fullhw bool) ffmpeg.VideoFilter {
	return args
}

// Switch scaler
func (c *Codec_hardware_m264) HwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter {
	var template string

	template = "scale_vt=$value"

	// BUG: scale_vt doesn't call ff_scale_adjust_dimensions, thus cant accept negative size values
	return ffmpeg.VideoFilter(templateReplaceScale(sargs, template, match, vf, true))
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_m264) HwCodecMaxRes() (int, int) {
	return 0, 0
}

// Return if a hardware accelerated for HLS is available
func (c *Codec_hardware_m264) HwCodecHLSCompatible() bool {
	return true
}

// Return if a hardware accelerated codec for MP4 is available
func (c *Codec_hardware_m264) HwCodecMP4Compatible() bool {
	return true
}

// Return if a hardware accelerated codec for WebM is available
func (c *Codec_hardware_m264) HwCodecWEBMCompatible() bool {
	return false
}

var _ codec.Codec = (*Codec_hardware_m264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_m264)(nil)
