package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
)

type Codec_hardware_r264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_r264() *Codec_hardware_r264 {
	c := &Codec_hardware_r264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_r264) Name() string {
	return "H264 V4L2M2M"
}

func (f *Codec_hardware_r264) CodeName() string {
	return "h264_v4l2m2m"
}

func (f *Codec_hardware_r264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_r264) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_r264) hwApplyFullHWFilter(args VideoFilter, fullhw bool) VideoFilter {
	return args
}

// Switch scaler
func (c *Codec_hardware_r264) hwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) VideoFilter {
	return VideoFilter(sargs)
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_r264) HwCodecMaxRes() (int, int) {
	return 0, 0
}

var _ codec.Codec = (*Codec_hardware_r264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_r264)(nil)
