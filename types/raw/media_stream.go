package raw

import (
	"fmt"
	"strings"
)

type Flags int

const (
	AUTO_DETECT Flags = 1 << iota
	REQUIRED
	IGNORE
	NO_LATENCY
)

func (f Flags) String() string {
	names := []string{"AUTO_DETECT", "REQUIRED", "IGNORE", "NO_LATENCY"}
	var flags []string
	for i := 0; i < len(names); i++ {
		if f&(1<<i) != 0 {
			flags = append(flags, names[i])
		}
	}
	return strings.Join(flags, "|")
}

type MediaStream struct {
	AudioParameters  string
	VideoParameters  string
	MediaPath        string
	AudioPath        string
	AudioFlags       Flags
	VideoFlags       Flags
	FfmpegParameters string
	YtDlpParameters  string
	Headers          map[string]string
}

func NewMediaStream(mediaPath string, audioParameters string, videoParameters string, audioPath string, audioFlags Flags, videoFlags Flags, headers map[string]string, ffmpegParameters string, ytDlpParameters string) *MediaStream {
	return &MediaStream{
		AudioParameters:  audioParameters,
		VideoParameters:  videoParameters,
		MediaPath:        mediaPath,
		AudioPath:        audioPath,
		AudioFlags:       audioFlags,
		VideoFlags:       videoFlags,
		FfmpegParameters: ffmpegParameters,
		YtDlpParameters:  ytDlpParameters,
		Headers:          headers,
	}
}

func (m *MediaStream) CheckStream() {
	if m.VideoFlags&IGNORE == 0 {
		fmt.Println("Check video stream")
	}
	if m.AudioFlags&IGNORE == 0 {
		fmt.Println("Check audio stream")
	}
}
