package message_test

import (
	"nonnonstop/akariai/discord/action"
	actions "nonnonstop/akariai/discord/action/message"
)

func getTestActionParam() *actions.ActionParam {
	return &actions.ActionParam{
		GuildID:   "TestGuildID",
		ChannelID: "TestChannelID",
		MessageID: "TestMessageID",
		User: action.DiscordUser{
			ID:    "TestUserID",
			Name:  "TestUserName",
			IsBot: false,
		},
		Content:       "TestContent",
		ContentOrg:    "TestContentOrg",
		IsMentionToMe: false,
	}
}

type actionDummy struct {
	Result bool
	Called chan struct{}
}

func (a *actionDummy) RunAction(d action.Discord, p *actions.ActionParam) bool {
	if a.Called != nil {
		a.Called <- struct{}{}
	}
	return a.Result
}
