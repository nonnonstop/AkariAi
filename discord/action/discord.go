package action

type Discord interface {
	SendMessage(channelID, message string)
	SendWebhook(channelID, name, message, iconURL string, file *FileContent)
	AddReaction(channelID, messageID, emoji string)
	AddRoleByName(guildID, userID, roleName string)
	HasRoleNyName(guildID, userID, roleName string) bool
	SetOnline(online bool)
	SetNickName(guildID, userID, nickName string)
}

type DiscordUser struct {
	ID    string
	Name  string
	IsBot bool
}

type FileContent struct {
	Name        string
	ContentType string
	Content     string
}
