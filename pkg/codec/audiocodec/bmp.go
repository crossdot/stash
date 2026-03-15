package videocodec_software

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_software_bmp struct {
	BaseSoftwareVideoCodec
}

func NewCodec_software_bmp() *Codec_software_bmp {
	c := &Codec_software_bmp{}
	c.BaseSoftwareVideoCodec.self = c
	return c
}

func (f *Codec_software_bmp) Name() string {
	return "BMP"
}

func (f *Codec_software_bmp) CodeName() string {
	return "bmp"
}

func (c *Codec_software_bmp) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	return args
}

var _ codec.Codec = (*Codec_software_bmp)(nil)
