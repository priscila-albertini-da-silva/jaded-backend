package models

import (
	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model
	Name     string
	Code     string
	Quantity int64
}
