package reaction

import "nonnonstop/akariai/discord/action"

type CheckBot struct {
	Target TargetUser
	Action Action
}

func (a *CheckBot) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if !p.getUser(a.Target).IsBot {
		return false
	}
	return a.Action.RunAction(d, p)
}
