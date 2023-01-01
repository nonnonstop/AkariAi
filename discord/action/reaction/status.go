package reaction

import "nonnonstop/akariai/discord/action"

type ActionStatus struct {
	Online bool
}

func (a *ActionStatus) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	d.SetOnline(a.Online)
	return false
}
