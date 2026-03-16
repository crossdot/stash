package videocodec_software

import "github.com/stashapp/stash/pkg/ffmpeg/codec"

type BaseSoftwareVideoCodec struct {
	self codec.VideoCodec
}

func (b *BaseSoftwareVideoCodec) Args() []string {
	return []string{"-c:v", string(b.self.CodeName())}
}
