package videocodec_software

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_software_libwebp struct {
	BaseSoftwareVideoCodec
}

func NewCodec_software_libwebp() *Codec_software_libwebp {
	c := &Codec_software_libwebp{}
	c.BaseSoftwareVideoCodec.self = c
	return c
}

func (f *Codec_software_libwebp) Name() string {
	return "WebP"
}

func (f *Codec_software_libwebp) CodeName() string {
	return "libwebp"
}

func (c *Codec_software_libwebp) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	return args
}

var _ codec.Codec = (*Codec_software_libwebp)(nil)
