package message

import (
	"nonnonstop/akariai/discord/action"
	"regexp"
)

type CheckRegex struct {
	Regex  *regexp.Regexp
	Action Action
}

func (a *CheckRegex) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if !a.Regex.MatchString(p.Content) {
		return false
	}
	return a.Action.RunAction(d, p)
}
