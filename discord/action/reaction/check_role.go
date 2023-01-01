package reaction

import "nonnonstop/akariai/discord/action"

type CheckRole struct {
	Name   string
	Target TargetUser
	Action Action
}

func (a *CheckRole) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if p.GuildID == "" {
		return false
	}
	if !d.HasRoleNyName(p.GuildID, p.getUser(a.Target).ID, a.Name) {
		return false
	}
	return a.Action.RunAction(d, p)
}
