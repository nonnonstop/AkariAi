package message

import "nonnonstop/akariai/discord/action"

type ActionInterrupt struct{}

func (a *ActionInterrupt) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	return true
}
