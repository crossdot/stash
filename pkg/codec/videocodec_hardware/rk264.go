package videocodec_hardware

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
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

// Returns the max resolution for a given codec, or a default
func (c *Codec_hardware_rk264) HwCodecMaxRes() (int, int) {
	return 8192, 8192
}

// Initialise a video filter for HW encoding
func (c *Codec_hardware_rk264) hwFilterInit(fullhw bool) VideoFilter {
	var videoFilter VideoFilter
	// For Rockchip full-hw, do NOT pre-map to rkrga here. scale_rkrga can
	// consume DRM_PRIME frames directly when filter_hw_device is set.
	// For non-fullhw, keep a sane software format.
	if !fullhw {
		videoFilter = videoFilter.Append("format=nv12")
		videoFilter = videoFilter.Append("hwupload")
	}
	return videoFilter
}

var _ codec.Codec = (*Codec_hardware_rk264)(nil)
var _ codec.HardwareCodec = (*Codec_hardware_rk264)(nil)
