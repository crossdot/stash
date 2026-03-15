package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
)

type Codec_hardware_a264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_a264() *Codec_hardware_a264 {
	c := &Codec_hardware_a264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (c *Codec_hardware_a264) Name() string {
	return "H264 Advanced Media Framework (AMF)"
}

func (c *Codec_hardware_a264) CodeName() string {
	return "h264_amf"
}

func (c *Codec_hardware_a264) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	args = append(args,
		"-quality", "speed",
	)
	return args
}

func (f *Codec_hardware_a264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_a264) HwFilterInit(fullhw bool) ffmpeg.VideoFilter {
	var videoFilter ffmpeg.VideoFilter
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_a264) HwApplyFullHWFilter(args ffmpeg.VideoFilter, fullhw bool) ffmpeg.VideoFilter {
	return args
}

// Switch scaler
func (c *Codec_hardware_a264) HwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter {
	return ffmpeg.VideoFilter(sargs)
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_a264) HwCodecMaxRes() (int, int) {
	return 0, 0
}

// Return if a hardware accelerated for HLS is available
func (c *Codec_hardware_a264) HwCodecHLSCompatible() bool {
	return false
}

// Return if a hardware accelerated codec for MP4 is available
func (c *Codec_hardware_a264) HwCodecMP4Compatible() bool {
	return false
}

// Return if a hardware accelerated codec for WebM is available
func (c *Codec_hardware_a264) HwCodecWEBMCompatible() bool {
	return false
}

var _ codec.Codec = (*Codec_hardware_a264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_a264)(nil)
