package message

import "nonnonstop/akariai/discord/action"

type ActionParam struct {
	GuildID       string
	ChannelID     string
	MessageID     string
	User          action.DiscordUser
	Content       string
	ContentOrg    string
	IsMentionToMe bool
}

type Action interface {
	RunAction(d action.Discord, p *ActionParam) bool
}
