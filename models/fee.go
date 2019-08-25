package models

import (
	"github.com/jinzhu/gorm"
)

type Fee struct {
	gorm.Model
	TotalValue float64
	ToValue    float64
	FeeValue   float64
}
