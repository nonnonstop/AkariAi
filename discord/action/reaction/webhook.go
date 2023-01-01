package reaction

import "nonnonstop/akariai/discord/action"

type ActionWebhook struct {
	Name    string
	Message string
	IconURL string
	File    *action.FileContent
}

func (a *ActionWebhook) RunAction(
	d action.Discord,
	p *ActionParam,
) bool {
	d.SendWebhook(p.ChannelID, a.Name, a.Message, a.IconURL, a.File)
	return false
}
