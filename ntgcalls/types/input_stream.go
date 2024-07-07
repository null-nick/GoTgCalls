package types

import "github.com/pytgcalls/gotgcalls/ntgcalls"

type MediaStream struct {
	AudioParameters AudioParameters
	VideoParameters VideoParameters
	StreamAudio     AudioStream
	StreamVideo     VideoStream
	InputMode       ntgcalls.InputMode
}
