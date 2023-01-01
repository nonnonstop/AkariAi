package reaction

import (
	"nonnonstop/akariai/discord/action"
)

type ActionParam struct {
	GuildID     string
	ChannelID   string
	MessageID   string
	User        action.DiscordUser
	MessageUser action.DiscordUser
	Emoji       string
}

type Action interface {
	RunAction(d action.Discord, p *ActionParam) bool
}

type TargetUser int

const (
	ReactingUser TargetUser = 0
	ReactedUser  TargetUser = 1
)

func (p *ActionParam) getUser(t TargetUser) *action.DiscordUser {
	switch t {
	case ReactingUser:
		return &p.User
	case ReactedUser:
		return &p.MessageUser
	default:
		return &action.DiscordUser{}
	}
}
