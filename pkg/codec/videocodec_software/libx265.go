package videocodec_software

import "github.com/stashapp/stash/pkg/codec"

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

var _ codec.Codec = (*Codec_software_libx265)(nil)
