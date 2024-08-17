package internal

import (
	"exchange/db"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

func RunasyncApi(DB db.DB, exchange Binance, pubsub PubSub) {
	log.Trace().Msg("Running async api")

	pubsub.Subscribe(GetBalanceEvent, func(m *nats.Msg) {
		response := GetBalanceResponse{
			Test:    exchange.test,
			Balance: exchange.GetBalance(),
		}

		pubsub.Publish(m.Reply, response)
	})

	pubsub.Subscribe(GetConfigsEvent, func(m *nats.Msg) {
		payload := GetConfigsResponse{
			Configs: DB.GetConfigs(),
		}

		pubsub.Publish(m.Reply, payload)
	})

	pubsub.Subscribe(GetPositionsEvent, func(m *nats.Msg) {
		response := GetPositionsResponse{
			Positions: DB.GetPositions(),
		}

		pubsub.Publish(m.Reply, response)
	})

	pubsub.Subscribe(GetTradesEvent, func(m *nats.Msg) {
		response := GetTradesResponse{
			Trades: DB.GetTrades(),
		}

		pubsub.Publish(m.Reply, response)
	})

	pubsub.Subscribe(GetStatsEvent, func(m *nats.Msg) {
		var response GetStatsResponse

		trades := DB.GetTrades()

		if len(trades) != 0 {
			stats := CalculateStats(trades)
			response = GetStatsResponse{&stats}
		}

		pubsub.Publish(m.Reply, response)
	})
}

func CalculateStats(trades []db.Trades) (stats Stats) {
	for _, trade := range trades {
		percentage := ((trade.Exit - trade.Entry) / trade.Entry) * 100
		price := trade.Quantity * trade.Exit
		amount := (percentage * price) / 100

		if amount > 0 {
			stats.Profit += amount
		} else {
			stats.Loss += -1 * amount
		}
	}

	stats.Total = stats.Profit + stats.Loss
	return
}
