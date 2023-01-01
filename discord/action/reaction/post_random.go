package reaction

import (
	"math/rand"
	"nonnonstop/akariai/discord/action"
)

type ActionPostRandom struct {
	Messages []string
	Mention  bool
	Target   TargetUser
}

func (a *ActionPostRandom) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	messages := a.Messages
	message := messages[rand.Intn(len(messages))]
	if a.Mention {
		message = "<@" + p.getUser(a.Target).ID + "> " + message
	}
	d.SendMessage(p.ChannelID, message)
	return false
}
