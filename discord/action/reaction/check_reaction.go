package reaction

import "nonnonstop/akariai/discord/action"

type CheckReaction struct {
	Emoji  string
	Action Action
}

func (a *CheckReaction) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if p.Emoji != a.Emoji {
		return false
	}
	return a.Action.RunAction(d, p)
}
