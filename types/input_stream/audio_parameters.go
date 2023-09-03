package input_stream

type AudioParameters struct {
	Bitrate  uint16 `default:"48000"`
	Channels uint8  `default:"2"`
}
