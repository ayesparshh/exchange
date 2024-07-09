package db

import (
	"time"

	"github.com/rs/zerolog/log"
)

type Trades struct {
	ID       uint      `gorm:"primaryKey;" json:"id"`
	Symbol   string    `gorm:"not null" json:"symbol"`
	Entry    float64   `gorm:"not null" json:"entry"`
	Exit     float64   `gorm:"not null" json:"exit"`
	Quantity float64   `gorm:"not null" json:"quantity"`
	Time     time.Time `gorm:"not null" json:"time"`
}

func (db DB) GetTrades() []Trades {
	var trades []Trades

	result := db.conn.Find(&trades)

	if result.Error != nil {
		log.Error().Err(result.Error).Msg("DB.Trades.GetTrades")
	}

	return trades
}
