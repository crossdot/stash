package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
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

func (c *Codec_hardware_n264) Name() string {
	return "H264 NVENC"
}

func (c *Codec_hardware_n264) CodeName() string {
	return "h264_nvenc"
}

func (c *Codec_hardware_n264) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	args = append(args,
		"-rc", "vbr",
		"-cq", "15",
	)
	return args
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
func (c *Codec_hardware_n264) HwFilterInit(fullhw bool) ffmpeg.VideoFilter {
	var videoFilter ffmpeg.VideoFilter
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload_cuda")
	}
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_n264) HwApplyFullHWFilter(args ffmpeg.VideoFilter, fullhw bool) ffmpeg.VideoFilter {
	if fullhw && f.version.Gteq(Version{major: 5}) { // Added in FFMpeg 5
		args = args.Append("scale_cuda=format=yuv420p")
	}
	return args
}

// Switch scaler
func (c *Codec_hardware_n264) HwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter {
	var template string

	template = "scale_cuda=$value"
	if fullhw && f.version.Gteq(Version{major: 5}) { // Added in FFMpeg 5
		template += ":format=yuv420p"
	}

	return ffmpeg.VideoFilter(templateReplaceScale(sargs, template, match, vf, false))
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_n264) HwCodecMaxRes() (int, int) {
	return 4096, 4096
}

// Return if a hardware accelerated for HLS is available
func (c *Codec_hardware_n264) HwCodecHLSCompatible() bool {
	return true
}

// Return if a hardware accelerated codec for MP4 is available
func (c *Codec_hardware_n264) HwCodecMP4Compatible() bool {
	return true
}

// Return if a hardware accelerated codec for WebM is available
func (c *Codec_hardware_n264) HwCodecWEBMCompatible() bool {
	return false
}

var _ codec.VideoCodec = (*Codec_hardware_n264)(nil)
var _ codec.HardwareVideoCodec = (*Codec_hardware_n264)(nil)
