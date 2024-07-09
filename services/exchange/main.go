package main

import (
	"os"

	"exchange/db"
	"exchange/internal"
	"exchange/utils"

	"github.com/joho/godotenv"
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
	wait := make(chan bool)

	NatsUrl := os.Getenv("NATS_URL")
	NatsUser := os.Getenv("NATS_USER")
	NatsPass := os.Getenv("NATS_PASS")

	DatabaseUrl := os.Getenv("DATABASE_URL")

	DB := db.New(DatabaseUrl)
	DB.Seed()

	pubsub := internal.NewPubSub(NatsUrl, NatsUser, NatsPass)
	defer pubsub.Close()

	BinanceKey := os.Getenv("BINANCE_KEY")
	BinanceSecret := os.Getenv("BINANCE_SECRET")
	Binancetestnet := utils.ParseBool(os.Getenv("BINANCE_TESTNET"))

	bex := internal.NewBinance(
		BinanceKey,
		BinanceSecret,
		Binancetestnet,
		pubsub,
		DB,
	)

	// go bex.Kline()

	internal.RunasyncApi(DB, bex, pubsub)

	<-wait
}
