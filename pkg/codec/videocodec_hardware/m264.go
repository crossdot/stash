package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_hardware_m264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_m264() *Codec_hardware_m264 {
	c := &Codec_hardware_m264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_m264) Name() string {
	return "H264 VideoToolbox"
}

func (f *Codec_hardware_m264) CodeName() string {
	return "h264_videotoolbox"
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

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_m264) HwCodecMaxRes() (int, int) {
	return 0, 0
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_m264) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload")
	}
	return videoFilter
}

var _ codec.Codec = (*Codec_hardware_m264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_m264)(nil)
