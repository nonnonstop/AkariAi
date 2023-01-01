package message

import "nonnonstop/akariai/discord/action"

type ActionPost struct {
	Message string
	Mention bool
}

func (a *ActionPost) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	message := a.Message
	if a.Mention {
		message = "<@" + p.User.ID + "> " + message
	}
	d.SendMessage(p.ChannelID, message)
	return false
}
