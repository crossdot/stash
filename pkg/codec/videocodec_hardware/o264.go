package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
)

type Codec_hardware_o264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_o264() *Codec_hardware_o264 {
	c := &Codec_hardware_o264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_o264) Name() string {
	return "H264 OMX"
}

func (f *Codec_hardware_o264) CodeName() string {
	return "h264_omx"
}

func (f *Codec_hardware_o264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_o264) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_o264) hwApplyFullHWFilter(args VideoFilter, fullhw bool) VideoFilter {
	return args
}

// Switch scaler
func (c *Codec_hardware_o264) hwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) VideoFilter {
	return VideoFilter(sargs)
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_o264) HwCodecMaxRes() (int, int) {
	return 0, 0
}

var _ codec.Codec = (*Codec_hardware_o264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_o264)(nil)
