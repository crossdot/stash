package videocodec_software

import "github.com/stashapp/stash/pkg/codec"

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

var _ codec.Codec = (*Codec_software_vpx)(nil)
