package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
)

type Codec_hardware_n264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_n264() *Codec_hardware_n264 {
	c := &Codec_hardware_n264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (f *Codec_hardware_n264) Name() string {
	return "H264 NVENC"
}

func (f *Codec_hardware_n264) CodeName() string {
	return "h264_nvenc"
}

func (f *Codec_hardware_n264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
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

// Initialise a video filter for HW encoding
func (c *Codec_hardware_n264) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload_cuda")
	}
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_n264) hwApplyFullHWFilter(args VideoFilter, fullhw bool) VideoFilter {
	if fullhw && f.version.Gteq(Version{major: 5}) { // Added in FFMpeg 5
		args = args.Append("scale_cuda=format=yuv420p")
	}
	return args
}

// Switch scaler
func (c *Codec_hardware_n264) hwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) VideoFilter {
	var template string

	template = "scale_cuda=$value"
	if fullhw && f.version.Gteq(Version{major: 5}) { // Added in FFMpeg 5
		template += ":format=yuv420p"
	}

	return VideoFilter(templateReplaceScale(sargs, template, match, vf, false))
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_n264) HwCodecMaxRes() (int, int) {
	return 4096, 4096
}

var _ codec.Codec = (*Codec_hardware_n264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_n264)(nil)
