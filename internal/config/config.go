package config

import (
	"github.com/kelseyhightower/envconfig"
)

type TelegramBotConfig struct {
	ApiToken string `envconfig:"TELEGRAM_BOT_API_TOKEN" required:"true"`
}

type Config struct {
	TelegramBotConfig TelegramBotConfig
}

func FromEnv() (Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return *cfg, err
	}

	return *cfg, nil
}
