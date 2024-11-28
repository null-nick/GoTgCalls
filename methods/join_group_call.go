package methods

import (
	"fmt"
	tg "github.com/amarnathcjd/gogram/telegram"
	ntgcalls "github.com/pytgcalls/gotgcalls/bindings"
	"github.com/pytgcalls/gotgcalls/types"
)

func JoinGroupCall(
	mtp *tg.Client,
	client *ntgcalls.Client,
	streamParams types.StreamParams,
	chatId any,
) {
	rawChannel, _ := mtp.ResolveUsername(chatId.(string))
	channel := rawChannel.(*tg.Channel)
	jsonParams, _ := client.CreateCall(channel.ID, ntgcalls.MediaDescription{
		Audio: &ntgcalls.AudioDescription{
			InputMode:     streamParams.InputMode,
			SampleRate:    streamParams.AudioParameters.Bitrate,
			ChannelCount:  streamParams.AudioParameters.Channels,
			BitsPerSample: 16,
			Input:         fmt.Sprintf("ffmpeg -i %s -f s16le -ac %d -ar %d -v quiet pipe:1", streamParams.AudioStream.Path, streamParams.AudioParameters.Channels, streamParams.AudioParameters.Bitrate),
		},
	})
	fullChatRaw, _ := mtp.ChannelsGetFullChannel(
		&tg.InputChannelObj{
			ChannelID:  channel.ID,
			AccessHash: channel.AccessHash,
		},
	)
	var haveVideo bool
	if streamParams.VideoStream != nil {
		haveVideo = true
	}
	fullChat := fullChatRaw.FullChat.(*tg.ChannelFull)
	callResRaw, _ := mtp.PhoneJoinGroupCall(
		&tg.PhoneJoinGroupCallParams{
			Muted:        false,
			VideoStopped: haveVideo,
			Call:         fullChat.Call,
			Params: &tg.DataJson{
				Data: jsonParams,
			},
			JoinAs: &tg.InputPeerUser{
				UserID:     streamParams.GroupCallConfig.JoinAs.UserID,
				AccessHash: streamParams.GroupCallConfig.JoinAs.AccessHash,
			},
		},
	)
	callRes := callResRaw.(*tg.UpdatesObj)
	for _, update := range callRes.Updates {
		switch update.(type) {
		case *tg.UpdateGroupCallConnection:
			phoneCall := update.(*tg.UpdateGroupCallConnection)
			_ = client.Connect(channel.ID, phoneCall.Params.Data)
		}
	}
}
