package videocodec_software

import (
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
)

type AudioCodec_libopus struct {
	BaseAudioCodec
}

func NewAudioCodec_libopus() *AudioCodec_libopus {
	c := &AudioCodec_libopus{}
	c.BaseAudioCodec.self = c
	return c
}

func (f *AudioCodec_libopus) CodeName() string {
	return "libopus"
}

var _ codec.AudioCodec = (*AudioCodec_libopus)(nil)
