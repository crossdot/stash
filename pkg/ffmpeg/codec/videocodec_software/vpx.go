package videocodec_software

import (
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
)

type Codec_software_vpx struct {
	BaseSoftwareVideoCodec
}

func NewCodec_software_vpx() *Codec_software_vpx {
	c := &Codec_software_vpx{}
	c.BaseSoftwareVideoCodec.self = c
	return c
}

func (f *Codec_software_vpx) Name() string {
	return "VPX-VP8"
}

func (f *Codec_software_vpx) CodeName() string {
	return "libvpx"
}

func (c *Codec_software_vpx) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	return args
}

var _ codec.VideoCodec = (*Codec_software_vpx)(nil)
