package config

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
)

// Конфигурация приложения
type Config struct {
	Environment string `long:"environment" env:"ENVIRONMENT" description:"App environment (develop, test, prod)" default:"develop"`

	LogLevel string `long:"log-level" description:"Log level: panic, fatal, warn, info, debug" env:"LOG_LEVEL" default:"warn"`

	MaxCpu int `long:"max-cpu" env:"MAX_CPU" description:"Max cpu usage (GOMAXPROC)" default:"0"`

	BotToken string `long:"bot-token" env:"BOT_TOKEN" description:"Telegram bot token" default:""`
}

func NewConfig() (*Config, error) {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default|flags.IgnoreUnknown)
	_, err := parser.Parse()
	if err != nil {
		parser.WriteHelp(log.Writer())
		return nil, fmt.Errorf("config parse failed: %w", err)
	}

	return &cfg, nil
}
