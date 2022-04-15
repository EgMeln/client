// Package model contain model of struct
package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

// GeneratedPrice struct that contain record info about new price
type GeneratedPrice struct {
	ID       uuid.UUID
	Ask      float64
	Bid      float64
	Symbol   string
	DoteTime string
}

// UnmarshalBinary unmarshal currency from byte
func (gen *GeneratedPrice) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, gen)
}

// Transaction struct that contain record info about transaction
type Transaction struct {
	ID         uuid.UUID
	PriceOpen  float64
	IsBay      bool
	Symbol     string
	PriceClose float64
	StopLoss   float64
	TakeProfit float64
}
