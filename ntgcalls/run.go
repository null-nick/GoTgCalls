package ntgcalls

func (c *Context) Run() {
	err := c.mtproto.Start()
	if err != nil {
		panic(err)
	}
	c.mtproto.Idle()
	return
}
