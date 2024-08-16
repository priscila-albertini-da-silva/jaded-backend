package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Stock struct {
	ID            uint
	Name          string
	Ticker        string
	Sector        string
	PurchasePrice decimal.Decimal
	CreatedAt     time.Time
}
