package models

import (
	"gorm.io/gorm"
)

type Foods struct {
	gorm.Model
	Name       string
	Price      string
	CategoryID uint
	Category   Category `gorm:"foreignKey:CategoryID"`
}
