package reaction

import "nonnonstop/akariai/discord/action"

type ActionRole struct {
	Name   string
	Target TargetUser
}

func (a *ActionRole) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if p.GuildID == "" {
		return false
	}
	d.AddRoleByName(p.GuildID, p.getUser(a.Target).ID, a.Name)
	return false
}
