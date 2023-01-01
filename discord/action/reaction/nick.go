package reaction

import "nonnonstop/akariai/discord/action"

type ActionNick struct {
	Name   string
	Target TargetUser
}

func (a *ActionNick) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	d.SetNickName(p.GuildID, p.getUser(a.Target).ID, a.Name)
	return false
}
