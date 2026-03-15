package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
)

type Codec_hardware_v264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_v264() *Codec_hardware_v264 {
	c := &Codec_hardware_v264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (c *Codec_hardware_v264) Name() string {
	return "H264 VAAPI"
}

func (c *Codec_hardware_v264) CodeName() string {
	return "h264_vaapi"
}

func (c *Codec_hardware_v264) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	args = append(args,
		"-qp", "20",
	)
	return args
}

func (f *Codec_hardware_v264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_v264) HwFilterInit(fullhw bool) ffmpeg.VideoFilter {
	var videoFilter ffmpeg.VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload")
	}
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_v264) HwApplyFullHWFilter(args ffmpeg.VideoFilter, fullhw bool) ffmpeg.VideoFilter {
	if fullhw && f.version.Gteq(Version{major: 3, minor: 1}) { // Added in FFMpeg 3.1
		args = args.Append("scale_vaapi=format=nv12")
	}
	return args
}

// Switch scaler
func (c *Codec_hardware_v264) HwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter {
	var template string

	template = "scale_vaapi=$value"
	if fullhw && f.version.Gteq(Version{major: 3, minor: 1}) { // Added in FFMpeg 3.1
		template += ":format=nv12"
	}

	return ffmpeg.VideoFilter(templateReplaceScale(sargs, template, match, vf, false))
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_v264) HwCodecMaxRes() (int, int) {
	return 0, 0
}

// Return if a hardware accelerated for HLS is available
func (c *Codec_hardware_v264) HwCodecHLSCompatible() bool {
	return true
}

// Return if a hardware accelerated codec for MP4 is available
func (c *Codec_hardware_v264) HwCodecMP4Compatible() bool {
	return false
}

// Return if a hardware accelerated codec for WebM is available
func (c *Codec_hardware_v264) HwCodecWEBMCompatible() bool {
	return false
}

var _ codec.Codec = (*Codec_hardware_v264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_v264)(nil)
