package bindings

type Instance struct {
	uid       uint32
	exists    bool
	streamEnd []StreamEndCallback
}
