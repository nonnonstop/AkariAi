package message

import "nonnonstop/akariai/discord/action"

type ActionRunNinja struct {
	Actions []Action
}

func (a *ActionRunNinja) RunAction(
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
