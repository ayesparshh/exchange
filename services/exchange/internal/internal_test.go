package internal_test

import (
	"exchange/db"
	"exchange/internal"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStats(t *testing.T) {
	type testCase struct {
		trades         []db.Trades
		expectedProfit float64
		expectedLoss   float64
	}

	var testCases = []testCase{
		{
			[]db.Trades{
				{ID: 1, Symbol: "ETHUSDT", Entry: 2700, Exit: 2710, Quantity: 0.0045, Time: time.Now()},
			},
			0.04516666666666667,
			0,
		},
		{
			[]db.Trades{
				{ID: 1, Symbol: "BTCUSDT", Entry: 42500, Exit: 42800, Quantity: 0.004675, Time: time.Now()},
			},
			1.4124,
			0,
		},
		{
			[]db.Trades{
				{ID: 1, Symbol: "SOLUSDT", Entry: 80, Exit: 84, Quantity: 1.5, Time: time.Now()},
			},
			6.3,
			0,
		},
		{
			[]db.Trades{
				{ID: 1, Symbol: "ADAUSDT", Entry: 0.827, Exit: 0.728, Quantity: 20, Time: time.Now()},
			},
			0,
			1.7429746070133008,
		},
	}

	for _, testCase := range testCases {
		stats := internal.CalculateStats(testCase.trades)

		assert.Equal(t, stats.Profit, testCase.expectedProfit)
		assert.Equal(t, stats.Loss, testCase.expectedLoss)
	}
}
