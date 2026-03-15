package videocodec_software

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_software_a264 struct {
	BaseSoftwareVideoCodec
}

func NewCodec_software_a264() *Codec_software_a264 {
	c := &Codec_software_a264{}
	c.BaseSoftwareVideoCodec.self = c
	return c
}

func (f *Codec_software_a264) Name() string {
	return "x264"
}

func (f *Codec_software_a264) CodeName() string {
	return "libx264"
}

func (c *Codec_software_a264) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	args = append(args,
		"-pix_fmt", "yuv420p",
		"-preset", "veryfast",
		"-crf", "25",
		"-sc_threshold", "0",
	)
	return args
}

var _ codec.Codec = (*Codec_software_a264)(nil)
