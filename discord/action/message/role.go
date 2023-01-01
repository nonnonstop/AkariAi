package message

import "nonnonstop/akariai/discord/action"

type ActionRole struct {
	Name string
}

func (a *ActionRole) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	if p.GuildID == "" {
		return false
	}
	d.AddRoleByName(p.GuildID, p.User.ID, a.Name)
	return false
}
