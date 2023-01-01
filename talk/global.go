package talk

import "go.uber.org/zap"

var globalClient *Client

func InitGlobal(
	config *Config,
	logger *zap.SugaredLogger,
) (*Client, error) {
	talk, err := New(config, logger)
	globalClient = talk
	return talk, err
}

func Talk(message, userID, userName string) string {
	return globalClient.Talk(message, userID, userName)
}
