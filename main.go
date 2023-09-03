package main

//#cgo LDFLAGS: -L . -lntgcalls -Wl,-rpath=./
import "C"
import (
	"github.com/gotd/td/telegram"
	"github.com/pytgcalls/gotgcalls/methods/stream"
	"github.com/pytgcalls/gotgcalls/types/input_stream"
)

func main() {
	client := telegram.NewClient(0, "", telegram.Options{})
	stream.JoinGroupCall(
		client.API(),
		0,
		input_stream.InputStream{
			StreamAudio: input_stream.AudioStream{
				Path: "",
				Parameters: input_stream.AudioParameters{
					Bitrate:  0,
					Channels: 0,
				},
			},
			StreamVideo: input_stream.VideoStream{
				Path: "",
				Parameters: input_stream.VideoParameters{
					Width:  0,
					Height: 0,
				},
			},
			InputMode: 0,
			LipSync:   false,
		},
		"",
		nil,
	)
}
