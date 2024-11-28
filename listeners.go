package ntgcalls

import ntgcalls "github.com/pytgcalls/gotgcalls/bindings"

func (ctx *Context) OnStreamEnd(handler ntgcalls.StreamEndCallback) {
	ctx.ntgcalls.OnStreamEnd(handler)
}
