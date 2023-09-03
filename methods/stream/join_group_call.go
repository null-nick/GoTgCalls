package stream

import (
	"context"
	"fmt"
	"github.com/gotd/td/tg"
	"github.com/pytgcalls/gotgcalls/ntgcalls"
	"github.com/pytgcalls/gotgcalls/types/input_stream"
)

func JoinGroupCall(client *tg.Client, chatId int64, stream input_stream.InputStream, inviteHash string, joinAs tg.InputPeerClass) {
	call := ntgcalls.NTgCalls()
	defer call.Free()
	res, err := call.CreateCall(chatId, ntgcalls.MediaDescription{
		Audio: &ntgcalls.AudioDescription{
			InputMode:     stream.InputMode,
			SampleRate:    stream.StreamAudio.Parameters.Bitrate,
			ChannelCount:  stream.StreamAudio.Parameters.Channels,
			Input:         "ffmpeg -i https://docs.evostream.com/sample_content/assets/sintel1m720p.mp4 -f s16le -ac 2 -ar 48k pipe:1",
			BitsPerSample: 16,
		},
	})

	fullChat, err := client.MessagesGetFullChat(context.Background(), chatId)
	if err != nil {
		panic(err)
		return
	}
	calls, _ := fullChat.FullChat.GetCall()
	groupCall, err := client.PhoneJoinGroupCall(
		context.Background(),
		&tg.PhoneJoinGroupCallRequest{
			Call:       calls,
			JoinAs:     joinAs,
			Params:     tg.DataJSON{Data: res},
			InviteHash: inviteHash,
		},
	)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("WORKS", res, err.Error())
	fmt.Println(groupCall, err.Error())

	call.OnStreamEnd(func(chatId int64, streamType ntgcalls.StreamType) {
		fmt.Println(chatId)
	})
}
