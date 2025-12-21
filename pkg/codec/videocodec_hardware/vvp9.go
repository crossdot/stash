package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_vvp9 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_vvp9) Name() string {
	return "VP9 VAAPI"
}

func (f *Codec_hardware_vvp9) CodeName() string {
	return "vp9_vaapi"
}

var _ codec.Codec = (*Codec_hardware_vvp9)(nil)
