package message

import "nonnonstop/akariai/discord/action"

type ActionNick struct {
	Name string
}

func (a *ActionNick) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	d.SetNickName(p.GuildID, p.User.ID, a.Name)
	return false
}
