package videocodec_software

import (
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
)

type AudioCodec_aac struct {
	BaseAudioCodec
}

func NewAudioCodec_aac() *AudioCodec_aac {
	c := &AudioCodec_aac{}
	c.BaseAudioCodec.self = c
	return c
}

func (f *AudioCodec_aac) CodeName() string {
	return "aac"
}

var _ codec.AudioCodec = (*AudioCodec_aac)(nil)
