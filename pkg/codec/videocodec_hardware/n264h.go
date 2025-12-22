package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_hardware_n264h struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_n264h() *Codec_hardware_n264h {
	c := &Codec_hardware_n264h{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_n264h) Name() string {
	return "H264 NVENC HQ profile"
}

func (f *Codec_hardware_n264h) CodeName() string {
	return "h264_nvenc"
}

func (f *Codec_hardware_n264h) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	args = append(args, "-hwaccel_device")
	args = append(args, "0")
	if fullhw {
		args = append(args, "-threads")
		args = append(args, "1")
		args = append(args, "-hwaccel")
		args = append(args, "cuda")
		args = append(args, "-hwaccel_output_format")
		args = append(args, "cuda")
	}
	return args
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_n264h) HwCodecMaxRes() (int, int) {
	return 4096, 4096
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_n264h) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload_cuda")
	}
	return videoFilter
}

var _ codec.Codec = (*Codec_hardware_n264h)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_n264h)(nil)
