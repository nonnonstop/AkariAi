package discord

import (
	"nonnonstop/akariai/discord/action"
	actions "nonnonstop/akariai/discord/action/reaction"

	"github.com/bwmarrin/discordgo"
)

func (d *Discord) onMessageReactionAdd(
	s *discordgo.Session,
	r *discordgo.MessageReactionAdd,
) {
	// Ignore if sleeping
	if d.status.Status == string(discordgo.StatusInvisible) {
		return
	}

	d.UpdateStatusToOnline()

	message, err := d.session.ChannelMessage(r.ChannelID, r.MessageID)
	if err != nil {
		d.logger.Errorln("Failed to get message", err)
	}
	reaction := &actions.ActionParam{
		GuildID:   r.GuildID,
		ChannelID: r.ChannelID,
		MessageID: r.MessageID,
		User: action.DiscordUser{
			ID:    r.UserID,
			Name:  r.Member.User.Username,
			IsBot: r.Member.User.Bot,
		},
		MessageUser: action.DiscordUser{
			ID:    message.Author.ID,
			Name:  message.Author.Username,
			IsBot: message.Author.Bot,
		},
		Emoji: r.Emoji.Name,
	}
	reactionActions.RunAction(d, reaction)
}

var reactionActions = &actions.ActionRun{
	Actions: []actions.Action{
		&actions.CheckDM{
			Action: &actions.ActionInterrupt{},
		},
		&actions.CheckBot{
			Target: actions.ReactingUser,
			Action: &actions.ActionInterrupt{},
		},
		&actions.CheckReaction{
			Emoji: "💩",
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.CheckRole{
						Name:   "うんこ",
						Target: actions.ReactingUser,
						Action: &actions.ActionRun{
							Actions: []actions.Action{
								&actions.ActionRole{
									Name:   "うんこ",
									Target: actions.ReactedUser,
								},
								&actions.ActionNick{
									Name: "うんこ",
								},
							},
						},
					},
					&actions.ActionInterrupt{},
				},
			},
		},
	},
}
