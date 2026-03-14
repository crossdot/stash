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
}

type HardwareCodec interface {
	// Tests all (given) hardware codec's
	InitHWSupport(context.Context) bool
	// Prepend input for hardware encoding only
	HwDeviceInit(ffmpeg.Args, bool) ffmpeg.Args
	// Initialise a video filter for HW encoding
	HwFilterInit(fullhw bool) VideoFilter
	// Apply format switching if applicable
	hwApplyFullHWFilter(args VideoFilter, fullhw bool) VideoFilter
	// Switch scaler
	hwApplyScaleTemplate(sargs string, match []int, vf *models.VideoFile, fullhw bool) VideoFilter
	// Returns the max resolution for a given codec, or a default
	HwCodecMaxRes() (int, int)
	// Return if a hardware accelerated for HLS is available
	hwCodecHLSCompatible() bool
	// Return if a hardware accelerated codec for MP4 is available
	hwCodecMP4Compatible() bool
	// Return if a hardware accelerated codec for WebM is available
	hwCodecWEBMCompatible() bool
}
