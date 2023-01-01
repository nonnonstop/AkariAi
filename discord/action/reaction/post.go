package reaction

import "nonnonstop/akariai/discord/action"

type ActionPost struct {
	Message string
	Mention bool
	Target  TargetUser
}

func (a *ActionPost) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	message := a.Message
	if a.Mention {
		message = "<@" + p.getUser(a.Target).ID + "> " + message
	}
	d.SendMessage(p.ChannelID, message)
	return false
}
