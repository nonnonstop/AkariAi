package reaction

import "nonnonstop/akariai/discord/action"

type ActionAsync struct {
	Action Action
}

func (a *ActionAsync) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	go func() {
		a.Action.RunAction(d, p)
	}()
	return false
}
