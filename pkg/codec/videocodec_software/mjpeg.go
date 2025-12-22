package videocodec_software

import "github.com/stashapp/stash/pkg/codec"

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

var _ codec.Codec = (*Codec_software_mjpeg)(nil)
