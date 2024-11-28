package types

type AudioParameters struct {
	Bitrate  uint32 `default:"48000"`
	Channels uint8  `default:"2"`
}
