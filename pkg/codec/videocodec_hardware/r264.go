package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_r264 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_r264) Name() string {
	return "H264 V4L2M2M"
}

func (f *Codec_hardware_r264) CodeName() string {
	return "h264_v4l2m2m"
}

var _ codec.Codec = (*Codec_hardware_r264)(nil)
