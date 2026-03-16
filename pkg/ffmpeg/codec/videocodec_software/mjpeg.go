package videocodec_software

import (
	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/ffmpeg/codec"
)

type Codec_software_mjpeg struct {
	BaseSoftwareVideoCodec
}

func NewCodec_software_mjpeg() *Codec_software_mjpeg {
	c := &Codec_software_mjpeg{}
	c.BaseSoftwareVideoCodec.self = c
	return c
}

func (f *Codec_software_mjpeg) Name() string {
	return "Jpeg"
}

func (f *Codec_software_mjpeg) CodeName() string {
	return "mjpeg"
}

func (c *Codec_software_mjpeg) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	return args
}

var _ codec.VideoCodec = (*Codec_software_mjpeg)(nil)
