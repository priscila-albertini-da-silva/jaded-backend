package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Dividend struct {
	ID          int64
	Description string
	Stock       Stock `gorm:"foreignKey:stock_id"`
	Value       decimal.Decimal
	PaymentDate time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
