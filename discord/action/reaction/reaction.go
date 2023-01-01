package reaction

import "nonnonstop/akariai/discord/action"

type ActionReacton struct {
	Emoji string
}

func (a *ActionReacton) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	d.AddReaction(p.ChannelID, p.MessageID, a.Emoji)
	return false
}
