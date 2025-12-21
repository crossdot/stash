package videocodec_hardware

import "github.com/stashapp/stash/pkg/codec"

type Codec_hardware_m264 struct {
	BaseHardwareVideoCodec
}

func (f *Codec_hardware_m264) Name() string {
	return "H264 VideoToolbox"
}

func (f *Codec_hardware_m264) CodeName() string {
	return "h264_videotoolbox"
}

var _ codec.Codec = (*Codec_hardware_m264)(nil)
