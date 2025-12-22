package codec

import (
	"context"

	"github.com/stashapp/stash/pkg/ffmpeg"
)

type Encoder interface {
	Encode([]byte) ([]byte, error)
}

type Decoder interface {
	Decode([]byte) ([]byte, error)
}

type Codec interface {
	// Encoder
	// Decoder
	Name() string
	CodeName() string
}

type HardwareCodec interface {
	InitHWSupport(context.Context) bool
	HwDeviceInit(ffmpeg.Args, bool) ffmpeg.Args
	HwFilterInit(fullhw bool) VideoFilter
	hwApplyFullHWFilter(args VideoFilter, fullhw bool) VideoFilter
	HwCodecMaxRes() (int, int)
}
