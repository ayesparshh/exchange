package main

import (
	"notification/internal"
	"os"
	"strconv"

	"github.com/joho/godotenv" //extra
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn().Err(err).Msg("No .env file found")
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	NatsUrl := os.Getenv("NATS_URL")
	NatsUser := os.Getenv("NATS_USER")
	NatsPass := os.Getenv("NATS_PASS")

	pubsub := internal.NewPubSub(NatsUrl, NatsUser, NatsPass)
	defer pubsub.Close()

	TelegramApiToken := os.Getenv("TELEGRAM_API_TOKEN")
	TelegramChatId := os.Getenv("TELEGRAM_CHAT_ID")

	// TelegramChatId = utils.ParseInt(TelegramChatId)

	value, _ := strconv.Atoi(TelegramChatId)

	telegram := internal.NewTelegramBot(TelegramApiToken, int64(value), pubsub)

	internal.RunAsyncApi(telegram, pubsub)
	telegram.ListenForCommands()
}
