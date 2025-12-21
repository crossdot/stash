package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_v264 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_v264) Name() string {
	return "H264 VAAPI"
}

func (f *Codec_hardware_v264) CodeName() string {
	return "h264_vaapi"
}

var _ codec.Codec = (*Codec_hardware_v264)(nil)
