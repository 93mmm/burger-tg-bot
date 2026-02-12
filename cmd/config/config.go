package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
)

type Config struct {
	Environment string `env:"ENVIRONMENT" env-default:"develop"`
	LogLevel    string `env:"LOG_LEVEL" env-default:"warn"`
	MaxCpu      int    `env:"MAX_CPU" env-default:"0"`
	BotToken    string `env:"BOT_TOKEN" env-required:"true"`

	Messages messages
}

type messages struct {
	DailyLink    string `env:"DAILY_LINK" env-required:"true"`
	GitMrURL     string `env:"GIT_MR_URL" env-required:"true"`
	DitGifID     string `env:"DIT_GIF_FILE_ID" env-required:"true"`
	GroupMembers string `env:"GROUP_MEMBERS" env-required:"true"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, errors.Wrap(err, "config error")
	}

	return &cfg, nil
}
