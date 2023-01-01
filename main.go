package main

import (
	"log"
	"math/rand"
	"nonnonstop/akariai/discord"
	"nonnonstop/akariai/talk"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Init logger
	logger, err := initLogger()
	if err != nil {
		log.Fatalln("Failed to create logger: ", err)
	}
	defer logger.Sync()

	logger.Infoln("Starting application...")

	// Init rand
	rand.Seed(time.Now().UnixNano())

	// Load config
	configPath := os.Getenv("APP_CONFIG")
	if configPath == "" {
		configPath = "config.json"
	}
	config, err := loadConfig(configPath)
	if err != nil {
		logger.Fatalln("Failed to create discord client: ", err)
	}

	// Init talk
	_, err = talk.InitGlobal(&config.Talk, logger.Named("talk"))
	if err != nil {
		logger.Fatalln("Failed to create talk client: ", err)
	}

	// Init discord
	discord, err := discord.New(&config.Discord, logger.Named("discord"))
	if err != nil {
		logger.Fatalln("Failed to create discord client: ", err)
	}
	defer discord.Close()

	// Wait forever
	stopBot := make(chan os.Signal, 1)
	signal.Notify(stopBot, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stopBot

	logger.Infoln("Stopping application...")
}
