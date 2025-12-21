package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_vvpx struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_vvpx) Name() string {
	return "VP8 VAAPI"
}

func (f *Codec_hardware_vvpx) CodeName() string {
	return "vp8_vaapi"
}

var _ codec.Codec = (*Codec_hardware_vvpx)(nil)
