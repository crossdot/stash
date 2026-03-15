package codec

import (
	"context"

	"github.com/stashapp/stash/pkg/ffmpeg"
	"github.com/stashapp/stash/pkg/models"
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

	CodecInit() (args ffmpeg.Args)
}

type HardwareCodec interface {
	// Tests all (given) hardware codec's
	InitHWSupport(context.Context) bool
	// Prepend input for hardware encoding only
	HwDeviceInit(ffmpeg.Args, bool) ffmpeg.Args
	// Initialise a video filter for HW encoding
	HwFilterInit(fullhw bool) ffmpeg.VideoFilter
	// Apply format switching if applicable
	HwApplyFullHWFilter(args ffmpeg.VideoFilter, fullhw bool) ffmpeg.VideoFilter
	// Switch scaler
	HwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) ffmpeg.VideoFilter
	// Returns the max resolution for a given codec, or a default
	HwCodecMaxRes() (int, int)
	// Return if a hardware accelerated for HLS is available
	HwCodecHLSCompatible() bool
	// Return if a hardware accelerated codec for MP4 is available
	HwCodecMP4Compatible() bool
	// Return if a hardware accelerated codec for WebM is available
	HwCodecWEBMCompatible() bool
}

type AudioCodec interface {
}
