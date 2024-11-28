package ntgcalls

func (ctx *Context) Run() {
	err := ctx.mtproto.Start()
	if err != nil {
		panic(err)
	}
	ctx.mtproto.Idle()
	return
}
