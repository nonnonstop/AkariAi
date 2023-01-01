package message

import "nonnonstop/akariai/discord/action"

type ActionRun struct {
	Actions []Action
}

func (a *ActionRun) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	for _, action := range a.Actions {
		interrupt := action.RunAction(d, p)
		if interrupt {
			return true
		}
	}
	return false
}
