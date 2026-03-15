package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
)

type Codec_hardware_rk264 struct {
	BaseHardwareVideoCodec
}

func NewCodec_hardware_rk264() *Codec_hardware_rk264 {
	c := &Codec_hardware_rk264{}
	c.BaseHardwareVideoCodec.self = c
	return c
}

func (c *Codec_hardware_rk264) Name() string {
	return "H264 Rockchip MPP (rkmpp)"
}

func (c *Codec_hardware_rk264) CodeName() string {
	return "h264_rkmpp"
}

func (c *Codec_hardware_rk264) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	return args
}

func (c *Codec_hardware_rk264) HwDeviceInit(args ffmpeg.Args, fullhw bool) ffmpeg.Args {
	// Rockchip: always create rkmpp device and make it the filter device, so
	// scale_rkrga and subsequent hwupload/hwmap operate in the right context.
	args = append(args, "-init_hw_device")
	args = append(args, "rkmpp=rk")
	args = append(args, "-filter_hw_device")
	args = append(args, "rk")
	if fullhw {
		args = append(args, "-hwaccel")
		args = append(args, "rkmpp")
		args = append(args, "-hwaccel_output_format")
		args = append(args, "drm_prime")
	}
	return args
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_rk264) HwFilterInit(fullhw bool) ffmpeg.VideoFilter {
	var videoFilter ffmpeg.VideoFilter
	// For Rockchip full-hw, do NOT pre-map to rkrga here. scale_rkrga can
	// consume DRM_PRIME frames directly when filter_hw_device is set.
	// For non-fullhw, keep a sane software format.
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload")
	}
	return videoFilter
}

// Apply format switching if applicable
func (c *Codec_hardware_rk264) HwApplyFullHWFilter(args ffmpeg.VideoFilter, fullhw bool) ffmpeg.VideoFilter {
	// For Rockchip, no extra mapping here. If there is no scale filter,
	// leave frames in DRM_PRIME for the encoder.
	return args
}

// Switch scaler
func (c *Codec_hardware_rk264) HwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter {
	var template string

	// The original filter chain is a fallback for maximum compatibility:
	// "scale_rkrga=$value:format=nv12,hwdownload,format=nv12,hwupload"
	// It avoids hwmap(rkrga→rkmpp) failures (-38/-12) seen on some builds
	// by downloading the scaled frame to system RAM and re-uploading it.
	// The filter chain below uses a zero-copy approach, passing the hardware-scaled
	// frame directly to the encoder. This is more efficient but may be less stable.
	template = "scale_rkrga=$value"

	// Rockchip's scale_rkrga supports -1/-2; don't apply minus-one hack here.
	return ffmpeg.VideoFilter(templateReplaceScale(sargs, template, match, vf, false))
}

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_rk264) HwCodecMaxRes() (int, int) {
	return 8192, 8192
}

// Return if a hardware accelerated for HLS is available
func (c *Codec_hardware_rk264) HwCodecHLSCompatible() bool {
	return true
}

// Return if a hardware accelerated codec for MP4 is available
func (c *Codec_hardware_rk264) HwCodecMP4Compatible() bool {
	return true
}

// Return if a hardware accelerated codec for WebM is available
func (c *Codec_hardware_rk264) HwCodecWEBMCompatible() bool {
	return false
}

var _ codec.Codec = (*Codec_hardware_rk264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_rk264)(nil)
