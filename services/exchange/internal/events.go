package internal

import (
	"exchange/db"
	"time"
)

var CriticalErrorEvent string = "Event:CriticalError"

type CriticalErrorEventPayload struct {
	Error string `json:"error"`
}

var GetBalanceEvent string = "Event:Balance:Get"

type Balance struct {
	Asset  string  `json:"asset"`
	Amount float64 `json:"amount"`
}

type GetBalanceResponse struct {
	Test    bool      `json:"test"`
	Balance []Balance `json:"balance"`
}

var GetConfigsEvent string = "Event:Configs:Get"

type GetConfigsResponse struct {
	Configs []db.Configs `json:"configs"`
}

var GetStatsEvent string = "Event:Stats:Get"

type Stats struct {
	Profit float64 `json:"profit"`
	Loss   float64 `json:"loss"`
	Total  float64 `json:"total"`
}

type GetStatsResponse struct {
	Stats *Stats `json:"stats"`
}

var GetPositionsEvent string = "Event:Positions:Get"

type Positions struct {
	Symbol   string    `json:"symbol"`
	Price    float64   `json:"price"`
	Quantity float64   `json:"quantity"`
	Time     time.Time `json:"time"`
}

type GetPositionsResponse struct {
	Positions []db.Positions `json:"positions"`
}

var GetTradesEvent string = "Event:Trades:Get"

type GetTradesResponse struct {
	Trades []db.Trades `json:"trades"`
}
