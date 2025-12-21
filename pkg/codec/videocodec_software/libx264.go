package videocodec_software

import "github.com/stashapp/stash/pkg/codec"

type Codec_software_a264 struct {
	BaseSoftwareVideoCodec
}

func (f *Codec_software_a264) Name() string {
	return "x264"
}

func (f *Codec_software_a264) CodeName() string {
	return "libx264"
}

var _ codec.Codec = (*Codec_software_a264)(nil)
