package reaction

import (
	"nonnonstop/akariai/discord/action"
	"time"
)

type ActionWait struct {
	Duration time.Duration
}

func (a *ActionWait) RunAction(
	d action.Discord,
	m *ActionParam,
) bool {
	time.Sleep(a.Duration)
	return false
}
