package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
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

func (c *Codec_hardware_r264) Name() string {
	return "H264 V4L2M2M"
}

func (c *Codec_hardware_r264) CodeName() string {
	return "h264_v4l2m2m"
}

func (c *Codec_hardware_r264) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	return args
}

func (f *Codec_hardware_r264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_r264) HwFilterInit(fullhw bool) ffmpeg.VideoFilter {
	var videoFilter ffmpeg.VideoFilter
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_r264) HwApplyFullHWFilter(args ffmpeg.VideoFilter, fullhw bool) ffmpeg.VideoFilter {
	return args
}

// Switch scaler
func (c *Codec_hardware_r264) HwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter {
	return ffmpeg.VideoFilter(sargs)
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_r264) HwCodecMaxRes() (int, int) {
	return 0, 0
}

// Return if a hardware accelerated for HLS is available
func (c *Codec_hardware_r264) HwCodecHLSCompatible() bool {
	return true
}

// Return if a hardware accelerated codec for MP4 is available
func (c *Codec_hardware_r264) HwCodecMP4Compatible() bool {
	return false
}

// Return if a hardware accelerated codec for WebM is available
func (c *Codec_hardware_r264) HwCodecWEBMCompatible() bool {
	return false
}

var _ codec.VideoCodec = (*Codec_hardware_r264)(nil)
var _ codec.HardwareVideoCodec = (*Codec_hardware_r264)(nil)
