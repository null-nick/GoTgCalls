package ntgcalls

import (
	"github.com/amarnathcjd/gogram/telegram"
	ntgcalls "github.com/pytgcalls/gotgcalls/bindings"
)

type Context struct {
	ntgcalls *ntgcalls.Client
	mtproto  *telegram.Client
}
