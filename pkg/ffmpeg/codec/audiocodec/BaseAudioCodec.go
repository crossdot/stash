package videocodec_software

import "github.com/stashapp/stash/pkg/ffmpeg/codec"

type BaseAudioCodec struct {
	self codec.AudioCodec
}

func (b *BaseAudioCodec) Args() []string {
	return []string{"-c:a", string(b.self.CodeName())}
}
