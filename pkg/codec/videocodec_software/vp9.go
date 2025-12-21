package videocodec_software

import "github.com/stashapp/stash/pkg/codec"

type Codec_software_vp9 struct {
	BaseSoftwareVideoCodec
}

func (f *Codec_software_vp9) Name() string {
	return "VPX-VP9"
}

func (f *Codec_software_vp9) CodeName() string {
	return "libvpx-vp9"
}

var _ codec.Codec = (*Codec_software_vp9)(nil)
