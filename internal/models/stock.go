package models

import "time"

type Stock struct {
	ID        int64
	Name      string
	Code      string
	Quantity  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
