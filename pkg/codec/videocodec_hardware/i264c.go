package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_i264 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_i264) Name() string {
	return "H264 Intel Quick Sync Video (QSV)"
}

func (f *Codec_hardware_i264) CodeName() string {
	return "h264_qsv"
}

var _ codec.Codec = (*Codec_hardware_i264)(nil)
