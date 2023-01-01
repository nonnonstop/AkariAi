package discord

import (
	"nonnonstop/akariai/discord/action"
	actions "nonnonstop/akariai/discord/action/message"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (d *Discord) onMessageUpdate(
	s *discordgo.Session,
	mc *discordgo.MessageUpdate,
) {
	// Ignore if sleeping
	if d.status.Status == string(discordgo.StatusInvisible) {
		return
	}

	d.UpdateStatusToOnline()

	content := mc.Content

	// Check if mention to me
	mentionToMePrefix := s.State.User.Mention()
	isMentionToMe := strings.HasPrefix(content, mentionToMePrefix)
	if isMentionToMe {
		content = content[len(mentionToMePrefix):]
	}

	// Create message object
	m, err := d.session.ChannelMessage(mc.ChannelID, mc.Message.ID)
	if err != nil {
		d.logger.Errorln("Failed to get message: ", err)
	}
	message := &actions.ActionParam{
		GuildID:   mc.GuildID,
		ChannelID: mc.ChannelID,
		MessageID: mc.Message.ID,
		User: action.DiscordUser{
			ID:    m.Author.ID,
			Name:  m.Author.Username,
			IsBot: m.Author.Bot,
		},
		Content:       normalizeString(content),
		ContentOrg:    content,
		IsMentionToMe: isMentionToMe,
	}
	if message.Content == "" {
		return
	}
	messageEditActions.RunAction(d, message)
}

var messageEditActions = &actions.ActionRun{
	Actions: []actions.Action{
		&actions.CheckRegex{
			Regex: regexp.MustCompile(`(?:うん[ちこ]|UNKO|UNTI|UNCHI|大便|💩)`),
			Action: &actions.ActionRun{
				Actions: []actions.Action{
					&actions.ActionReacton{
						Emoji: "💩",
					},
					&actions.ActionRole{
						Name: "うんこ",
					},
					&actions.ActionNick{
						Name: "うんこ",
					},
				},
			},
		},
	},
}
