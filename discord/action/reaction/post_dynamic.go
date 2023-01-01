package reaction

import "nonnonstop/akariai/discord/action"

type ActionPostDynamic struct {
	Func func(action.Discord, *ActionParam) string
}

func (a *ActionPostDynamic) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	message := a.Func(d, p)
	d.SendMessage(p.ChannelID, message)
	return false
}
