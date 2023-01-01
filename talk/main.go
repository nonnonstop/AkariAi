package talk

import "go.uber.org/zap"

type Client struct {
	config *Config
	logger *zap.SugaredLogger
}

type Config struct {
	Token string
}

func New(
	config *Config,
	logger *zap.SugaredLogger,
) (*Client, error) {
	talk := &Client{
		config: config,
		logger: logger,
	}
	return talk, nil
}
