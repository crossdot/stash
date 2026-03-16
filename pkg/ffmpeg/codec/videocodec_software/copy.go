package videocodec_software

import (
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
)

type Codec_software_copy struct {
	BaseSoftwareVideoCodec
}

func NewCodec_software_copy() *Codec_software_copy {
	c := &Codec_software_copy{}
	c.BaseSoftwareVideoCodec.self = c
	return c
}

func (f *Codec_software_copy) Name() string {
	return "Copy"
}

func (f *Codec_software_copy) CodeName() string {
	return "copy"
}

func (c *Codec_software_copy) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	return args
}

var _ codec.VideoCodec = (*Codec_software_copy)(nil)
