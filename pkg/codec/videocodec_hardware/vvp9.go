package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_hardware_vvp9 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_vvp9() *Codec_hardware_vvp9 {
	c := &Codec_hardware_vvp9{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_vvp9) Name() string {
	return "VP9 VAAPI"
}

func (f *Codec_hardware_vvp9) CodeName() string {
	return "vp9_vaapi"
}

func (f *Codec_hardware_vvp9) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	args = append(args, "-vaapi_device")
	args = append(args, "/dev/dri/renderD128")
	if fullhw {
		args = append(args, "-hwaccel")
		args = append(args, "vaapi")
		args = append(args, "-hwaccel_output_format")
		args = append(args, "vaapi")
	}
	return args
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_vvp9) HwCodecMaxRes() (int, int) {
	return 0, 0
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_vvp9) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload")
	}
	return videoFilter
}

var _ codec.Codec = (*Codec_hardware_vvp9)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_vvp9)(nil)
