package videocodec_software

import "github.com/stashapp/stash/pkg/codec"

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

var _ codec.Codec = (*Codec_software_bmp)(nil)
