package message

import "nonnonstop/akariai/discord/action"

type CheckDM struct {
	Action Action
}

func (a *CheckDM) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if p.GuildID != "" {
		return false
	}
	return a.Action.RunAction(d, p)
}
