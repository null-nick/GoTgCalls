package types

import ntgcalls "github.com/pytgcalls/gotgcalls/bindings"

type MediaStream struct {
	AudioParameters AudioParameters
	VideoParameters VideoParameters
	StreamAudio     AudioStream
	StreamVideo     VideoStream
	InputMode       ntgcalls.InputMode
}
