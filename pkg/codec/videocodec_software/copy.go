package videocodec_software

import "github.com/stashapp/stash/pkg/codec"

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

var _ codec.Codec = (*Codec_software_copy)(nil)
