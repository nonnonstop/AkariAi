package message

import (
	"nonnonstop/akariai/discord/action"
	"strings"
)

type CheckContains struct {
	Substr string
	Action Action
}

func (a *CheckContains) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if !strings.Contains(p.Content, a.Substr) {
		return false
	}
	return a.Action.RunAction(d, p)
}
