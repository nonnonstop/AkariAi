package main

import (
	"encoding/json"
	"nonnonstop/akariai/discord"
	"nonnonstop/akariai/talk"
	"os"
)

type AppConfig struct {
	Discord discord.Config
	Talk    talk.Config
}

func loadConfig(path string) (*AppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &AppConfig{}
	err = json.Unmarshal(data, config)
	return config, err
}
