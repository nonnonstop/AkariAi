package message

import "nonnonstop/akariai/discord/action"

type CheckMention struct {
	Action Action
}

func (a *CheckMention) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if !p.IsMentionToMe {
		return false
	}
	return a.Action.RunAction(d, p)
}
