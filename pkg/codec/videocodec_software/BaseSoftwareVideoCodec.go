package videocodec_software

type BaseSoftwareVideoCodec struct{}

func (f *BaseSoftwareVideoCodec) InitHWSupport() bool {
	return true
}
