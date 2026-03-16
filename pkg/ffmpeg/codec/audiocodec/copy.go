package videocodec_software

import (
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
)

type AudioCodec_copy struct {
	BaseAudioCodec
}

func NewAudioCodec_copy() *AudioCodec_copy {
	c := &AudioCodec_copy{}
	c.BaseAudioCodec.self = c
	return c
}

func (f *AudioCodec_copy) Name() string {
	return "copy"
}

var _ codec.AudioCodec = (*AudioCodec_copy)(nil)
