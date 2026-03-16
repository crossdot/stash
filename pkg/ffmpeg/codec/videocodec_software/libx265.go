package videocodec_software

import (
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
)

type Codec_software_libx265 struct {
	BaseSoftwareVideoCodec
}

func NewCodec_software_libx265() *Codec_software_libx265 {
	c := &Codec_software_libx265{}
	c.BaseSoftwareVideoCodec.self = c
	return c
}

func (f *Codec_software_libx265) Name() string {
	return "x265"
}

func (f *Codec_software_libx265) CodeName() string {
	return "libx265"
}

func (c *Codec_software_libx265) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	return args
}

var _ codec.VideoCodec = (*Codec_software_libx265)(nil)
