package message

import "nonnonstop/akariai/discord/action"

type CheckBot struct {
	Action Action
}

func (a *CheckBot) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if !p.User.IsBot {
		return false
	}
	return a.Action.RunAction(d, p)
}
