package types

import (
	"github.com/pytgcalls/gotgcalls/ntgcalls/bindings"
)

type StreamParams struct {
	AudioParameters   AudioParameters
	VideoParameters   VideoParameters
	InputMode         bindings.InputMode
	PrivateCallConfig PrivateCallConfig
	GroupCallConfig   GroupCallConfig
	AudioStream       *AudioStream
	VideoStream       *VideoStream
}
