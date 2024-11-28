package bindings

type Client struct {
	uid       uint32
	exists    bool
	streamEnd []StreamEndCallback
}
