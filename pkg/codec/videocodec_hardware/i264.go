package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_hardware_i264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_i264() *Codec_hardware_i264 {
	c := &Codec_hardware_i264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_i264) Name() string {
	return "H264 Intel Quick Sync Video (QSV)"
}

func (f *Codec_hardware_i264) CodeName() string {
	return "h264_qsv"
}

func (f *Codec_hardware_i264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
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
func (c *Codec_hardware_i264) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("hwupload=extra_hw_frames=64")
		videoFilter = videoFilter.Append("format=qsv")
	}
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_i264) hwApplyFullHWFilter(args VideoFilter, fullhw bool) VideoFilter {
	if fullhw && f.version.Gteq(Version{major: 3, minor: 3}) { // Added in FFMpeg 3.3
		args = args.Append("scale_qsv=format=nv12")
	}
	return args
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_i264) HwCodecMaxRes() (int, int) {
	return 4096, 4096
}

var _ codec.Codec = (*Codec_hardware_i264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_i264)(nil)
