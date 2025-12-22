package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_hardware_vvpx struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_vvpx() *Codec_hardware_vvpx {
	c := &Codec_hardware_vvpx{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_vvpx) Name() string {
	return "VP8 VAAPI"
}

func (f *Codec_hardware_vvpx) CodeName() string {
	return "vp8_vaapi"
}

func (f *Codec_hardware_vvpx) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
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

// Initialise a video filter for HW encoding
func (c *Codec_hardware_vvpx) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_vvpx) hwApplyFullHWFilter(args VideoFilter, fullhw bool) VideoFilter {
	return args
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_vvpx) HwCodecMaxRes() (int, int) {
	return 0, 0
}

var _ codec.Codec = (*Codec_hardware_vvpx)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_vvpx)(nil)
