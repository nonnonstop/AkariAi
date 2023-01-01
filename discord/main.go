package discord

import (
	"bytes"
	"io/ioutil"
	"nonnonstop/akariai/discord/action"
	"path/filepath"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Discord struct {
	logger         *zap.SugaredLogger
	session        *discordgo.Session
	config         *Config
	status         discordgo.UpdateStatusData
	statusUpdateCh chan discordgo.Status
	statusOnlineCh chan struct{}
	statusReadyCh  chan struct{}
}

type Config struct {
	Token   string
	Webhook map[string]struct {
		ID    string
		Token string
	}
	Assets string
}

func New(
	config *Config,
	logger *zap.SugaredLogger,
) (*Discord, error) {
	s, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, err
	}
	s.Identify.Intents |= discordgo.IntentMessageContent
	d := Discord{
		logger:  logger,
		session: s,
		config:  config,
		status: discordgo.UpdateStatusData{
			Status: string(discordgo.StatusOnline),
		},
		statusUpdateCh: make(chan discordgo.Status),
		statusOnlineCh: make(chan struct{}),
		statusReadyCh:  make(chan struct{}),
	}
	s.AddHandler(d.onReady)
	s.AddHandler(d.onMessageCreate)
	s.AddHandler(d.onMessageUpdate)
	s.AddHandler(d.onMessageReactionAdd)
	d.startStatusThread()
	err = s.Open()
	if err != nil {
		return nil, err
	}
	return &d, err
}

func (d *Discord) Close() error {
	return d.session.Close()
}

func (d *Discord) SendMessage(channelID, message string) {
	_, err := d.session.ChannelMessageSend(channelID, message)
	if err != nil {
		d.logger.Errorln("Failed to send message: ", err)
	}
}

func (d *Discord) SendWebhook(channelID, name, message, iconURL string, file *action.FileContent) {
	c, ok := d.config.Webhook[channelID]
	if !ok {
		d.logger.Warnln("Webhook config not found: ", channelID)
		return
	}
	params := &discordgo.WebhookParams{
		Username:  name,
		Content:   message,
		AvatarURL: iconURL,
	}
	if file != nil {
		data, err := ioutil.ReadFile(filepath.Join(d.config.Assets, file.Content))
		if err != nil {
			d.logger.Errorln("Failed to read file: ", err)
		}
		params.Files = []*discordgo.File{
			{
				Name:        file.Name,
				ContentType: file.ContentType,
				Reader:      bytes.NewReader(data),
			},
		}
	}
	_, err := d.session.WebhookExecute(c.ID, c.Token, false, params)
	if err != nil {
		d.logger.Errorln("Failed to send webhook: ", err)
	}
}

func (d *Discord) AddReaction(channelID, messageID, emoji string) {
	err := d.session.MessageReactionAdd(channelID, messageID, emoji)
	if err != nil {
		d.logger.Errorln("Failed to send message: ", err)
	}
}

func (d *Discord) FindRoleByName(guildID, userID, roleName string) string {
	roles, err := d.session.GuildRoles(guildID)
	if err != nil {
		d.logger.Errorln("Failed to get roles: ", err)
		return ""
	}
	for _, role := range roles {
		if role.Name == roleName {
			return role.ID
		}
	}
	d.logger.Errorln("Failed to find role: ", err)
	return ""
}

func (d *Discord) AddRoleByName(guildID, userID, roleName string) {
	roleID := d.FindRoleByName(guildID, userID, roleName)
	if roleID == "" {
		return
	}
	err := d.session.GuildMemberRoleAdd(guildID, userID, roleID)
	if err != nil {
		d.logger.Errorln("Failed to add role: ", err)
	}
}

func (d *Discord) HasRoleNyName(guildID, userID, roleName string) bool {
	roleID := d.FindRoleByName(guildID, userID, roleName)
	if roleID == "" {
		return false
	}
	member, err := d.session.GuildMember(guildID, userID)
	if err != nil {
		d.logger.Errorln("Failed to get member info: ", err)
	}
	for _, memberRole := range member.Roles {
		if memberRole == roleID {
			return true
		}
	}
	return false
}

func (d *Discord) SetOnline(online bool) {
	if online {
		d.statusUpdateCh <- discordgo.StatusOnline
	} else {
		d.statusUpdateCh <- discordgo.StatusInvisible
	}
}

func (d *Discord) SetNickName(guildID, userID, nickName string) {
	err := d.session.GuildMemberNickname(guildID, userID, nickName)
	if err != nil {
		d.logger.Errorln("Failed to update nickname: ", err)
	}
}

func (d *Discord) UpdateStatusToOnline() {
	d.statusOnlineCh <- struct{}{}
}
