package codec

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
	InitHWSupport() bool
}
