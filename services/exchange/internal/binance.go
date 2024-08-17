package internal

import (
	"context"
	"errors"
	"exchange/db"
	"exchange/utils"

	"github.com/adshao/go-binance/v2"
	"github.com/rs/zerolog/log"
)

var ZeroBalance = 0.00000000

type Binance struct {
	client *binance.Client
	test   bool
	pubsub PubSub
	DB     db.DB
}

var ErrBaseAsset = errors.New("base asset for symbol not found")

func NewBinance(key, secret string, test bool, pubsub PubSub, DB db.DB) Binance {
	log.Trace().Str("type", "binance").Bool("test", test).Msg("Binance.Init")

	binance.UseTestnet = test
	client := binance.NewClient(key, secret)

	return Binance{client, test, pubsub, DB}
}

func (b Binance) GetAccount() *binance.Account {

	svc := b.client.NewGetAccountService()
	account, err := svc.Do(context.Background())

	if err != nil {
		log.Error().Err(err).Msg("Binance.UserInfo")
	}

	return account
}

func (b Binance) GetBalance() []Balance {
	acc := b.GetAccount()
	balances := []Balance{}

	for _, balance := range acc.Balances {
		asset := balance.Asset
		amt := utils.ParseFloat(balance.Free)

		if amt > ZeroBalance {
			b := Balance{asset, amt}
			balances = append(balances, b)
		}
	}

	return balances
}

func (b Binance) GetBalanceQuantity(symbol string) (float64, error) {
	info, err := b.client.NewExchangeInfoService().Symbol(symbol).Do(context.Background())

	if err != nil {
		log.Error().Str("symbol", symbol).Err(err).Msg("Binance.GetBalanceQuantity")
		return 0, err
	}

	balances := b.GetBalance()

	asset := info.Symbols[0].BaseAsset

	for _, balance := range balances {
		if balance.Asset == asset {
			return balance.Amount, nil
		}
	}

	log.Error().Str("symbol", symbol).Err(ErrBaseAsset).Msg("Binance.GetBalanceQuantity")
	b.pubsub.Publish(CriticalErrorEvent, CriticalErrorEventPayload{ErrBaseAsset.Error()})

	return 0, ErrBaseAsset
}
