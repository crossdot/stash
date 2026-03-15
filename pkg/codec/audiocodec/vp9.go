package videocodec_software

import (
	"github.com/stashapp/stash/pkg/codec"
	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Codec_software_vp9 struct {
	BaseSoftwareVideoCodec
}

func NewCodec_software_vp9() *Codec_software_vp9 {
	c := &Codec_software_vp9{}
	c.BaseSoftwareVideoCodec.self = c
	return c
}

func (f *Codec_software_vp9) Name() string {
	return "VPX-VP9"
}

func (f *Codec_software_vp9) CodeName() string {
	return "libvpx-vp9"
}

func (c *Codec_software_vp9) CodecInit() (args ffmpeg.Args) {
	args = args.VideoCodec(codec)
	args = append(args,
		"-pix_fmt", "yuv420p",
		"-deadline", "realtime",
		"-cpu-used", "5",
		"-row-mt", "1",
		"-crf", "30",
		"-b:v", "0",
	)
	return args
}

var _ codec.Codec = (*Codec_software_vp9)(nil)
