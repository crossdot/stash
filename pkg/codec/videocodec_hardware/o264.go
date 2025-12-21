package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_o264 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_o264) Name() string {
	return "H264 OMX"
}

func (f *Codec_hardware_o264) CodeName() string {
	return "h264_omx"
}

var _ codec.Codec = (*Codec_hardware_o264)(nil)
