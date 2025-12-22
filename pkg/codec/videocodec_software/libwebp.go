package videocodec_software

import "github.com/stashapp/stash/pkg/codec"

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

var _ codec.Codec = (*Codec_software_libwebp)(nil)
