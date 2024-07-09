package internal

import (
	"exchange/db"
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

type Stats struct {
	Profit float64 `json:"profit"`
	Loss   float64 `json:"loss"`
	Total  float64 `json:"total"`
}

type GetStatsResponse struct {
	Stats *Stats `json:"stats"`
}
