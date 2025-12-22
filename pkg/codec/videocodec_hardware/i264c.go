package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_hardware_i264c struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_i264c() *Codec_hardware_i264c {
	c := &Codec_hardware_i264c{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_i264c) Name() string {
	return "H264 Intel Quick Sync Video (QSV) Compatibility profile"
}

func (f *Codec_hardware_i264c) CodeName() string {
	return "h264_qsv"
}

func (f *Codec_hardware_i264c) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
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

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_i264c) HwCodecMaxRes() (int, int) {
	return 4096, 4096
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_i264c) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("hwupload=extra_hw_frames=64")
		videoFilter = videoFilter.Append("format=qsv")
	}
	return videoFilter
}

var _ codec.Codec = (*Codec_hardware_i264c)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_i264c)(nil)
