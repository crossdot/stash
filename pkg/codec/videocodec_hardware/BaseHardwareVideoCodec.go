package videocodec_hardware

type BaseHardwareVideoCodec struct{}

func (f *BaseHardwareVideoCodec) InitHWSupport() bool {
	return true
}
