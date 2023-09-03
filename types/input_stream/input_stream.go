package input_stream

import "github.com/pytgcalls/gotgcalls/ntgcalls"

type InputStream struct {
	StreamAudio AudioStream
	StreamVideo VideoStream
	InputMode   ntgcalls.InputMode
	LipSync     bool `default:"true"`
}
