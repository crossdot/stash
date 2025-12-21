package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_i264c struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_i264c) Name() string {
	return "H264 Intel Quick Sync Video (QSV) Compatibility profile"
}

func (f *Codec_hardware_i264c) CodeName() string {
	return "h264_qsv"
}

var _ codec.Codec = (*Codec_hardware_i264c)(nil)
