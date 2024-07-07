package ntgcalls

/*
#cgo LDFLAGS: -L/root/gotgcalls/ntgcalls/bindings/lib -lntgcalls -Wl,-rpath=/root/gotgcalls/ntgcalls/bindings/lib
*/
import "C"

import (
	"github.com/amarnathcjd/gogram/telegram"
	ntgcalls "github.com/pytgcalls/gotgcalls/ntgcalls/bindings"
)

func Client(apiId int32, apiHash string, sessionName string, testBackend bool) *Context {
	client := ntgcalls.NTgCalls()
	defer client.Free()
	mtproto, _ := telegram.NewClient(telegram.ClientConfig{
		AppID:    apiId,
		AppHash:  apiHash,
		TestMode: testBackend,
		Session:  sessionName,
	})

	return &Context{
		ntgcalls: client,
		mtproto:  mtproto,
	}
}
