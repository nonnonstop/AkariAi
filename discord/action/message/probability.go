package message

import (
	"math/rand"
	"nonnonstop/akariai/discord/action"
)

type ActionProbability struct {
	Probability int
	Action      Action
}

func (a *ActionProbability) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if rand.Intn(a.Probability) != 0 {
		return false
	}
	return a.Action.RunAction(d, p)
}
