package types

import tg "github.com/amarnathcjd/gogram/telegram"

type GroupCallConfig struct {
	InviteHash string `default:""`
	JoinAs     tg.InputPeerUser
	AutoStart  bool `default:"true"`
}
