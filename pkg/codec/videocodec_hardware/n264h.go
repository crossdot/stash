package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_n264h struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_n264h) Name() string {
	return "H264 NVENC HQ profile"
}

func (f *Codec_hardware_n264h) CodeName() string {
	return "h264_nvenc"
}

var _ codec.Codec = (*Codec_hardware_n264h)(nil)
