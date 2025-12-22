package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_hardware_a264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_a264() *Codec_hardware_a264 {
	c := &Codec_hardware_a264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_a264) Name() string {
	return "H264 Advanced Media Framework (AMF)"
}

func (f *Codec_hardware_a264) CodeName() string {
	return "h264_amf"
}

func (f *Codec_hardware_a264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_a264) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_a264) hwApplyFullHWFilter(args VideoFilter, fullhw bool) VideoFilter {
	return args
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_a264) HwCodecMaxRes() (int, int) {
	return 0, 0
}

var _ codec.Codec = (*Codec_hardware_a264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_a264)(nil)
