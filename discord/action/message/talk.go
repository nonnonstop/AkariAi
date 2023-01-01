package message

import (
	"nonnonstop/akariai/discord/action"
	"nonnonstop/akariai/talk"
)

type ActionTalk struct {
	Mention bool
}

func (a *ActionTalk) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	message := talk.Talk(p.ContentOrg, p.User.ID, p.User.Name)
	if a.Mention {
		message = "<@" + p.User.ID + "> " + message
	}
	d.SendMessage(p.ChannelID, message)
	return false
}
