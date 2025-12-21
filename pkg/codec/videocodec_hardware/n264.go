package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_n264 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_n264) Name() string {
	return "H264 NVENC"
}

func (f *Codec_hardware_n264) CodeName() string {
	return "h264_nvenc"
}

var _ codec.Codec = (*Codec_hardware_n264)(nil)
