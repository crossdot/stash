package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_ivp9 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_ivp9) Name() string {
	return "VP9 Intel Quick Sync Video (QSV)"
}

func (f *Codec_hardware_ivp9) CodeName() string {
	return "vp9_qsv"
}

var _ codec.Codec = (*Codec_hardware_ivp9)(nil)
