package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID        int64
	Stock     Stock
	Value     decimal.Decimal
	Quantity  float64
	Movement  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Movement string

const (
	BUY  Movement = "BUY"
	SELL Movement = "SELL"
)
