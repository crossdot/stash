package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_a264 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_a264) Name() string {
	return "H264 Advanced Media Framework (AMF)"
}

func (f *Codec_hardware_a264) CodeName() string {
	return "h264_amf"
}

var _ codec.Codec = (*Codec_hardware_a264)(nil)
