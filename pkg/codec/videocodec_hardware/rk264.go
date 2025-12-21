package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_rk264 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_rk264) Name() string {
	return "H264 Rockchip MPP (rkmpp)"
}

func (f *Codec_hardware_rk264) CodeName() string {
	return "h264_rkmpp"
}

var _ codec.Codec = (*Codec_hardware_rk264)(nil)
