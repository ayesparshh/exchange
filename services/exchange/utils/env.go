package utils

import "github.com/kelseyhightower/envconfig"

type Env struct {
	Binancetestnet bool   `envconfig:"Binance_TESTNET"`
	BinanceKey     string `envconfig:"Binance_KEY"`
	BinanceSecret  string `envconfig:"Binance_SECRET"`

	NatsUser string `envconfig:"NATS_USER"`
	NatsPass string `envconfig:"NATS_PASS"`
	NatsUrl  string `envconfig:"NATS_URL"`

	DatabaseUrl string `envconfig:"DATABASE_URL"`
}

func GetEnv() Env {
	var env Env
	envconfig.MustProcess("", &env)
	return env
}
