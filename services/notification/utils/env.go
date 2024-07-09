package utils

import "github.com/kelseyhightower/envconfig"

type Env struct {
	TelegramApiToken string `envconfig:"TELEGRAM_API_TOKEN"`
	TelegramChatId   int64  `envconfig:"TELEGRAM_CHAT_ID"`

	NatsUrl  string `envconfig:"NATS_URL"`
	NatsUser string `envconfig:"NATS_USER"`
	NatsPass string `envconfig:"NATS_PASS"`
}

func GetEnv() Env {
	var env Env
	envconfig.MustProcess("", &env)
	return env
}
