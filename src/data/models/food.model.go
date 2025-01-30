package models

import (
	"gorm.io/gorm"
)

type Foods struct {
	gorm.Model
	Name       string
	Price      string
	Category   Category `gorm:"foreignKey:CategoryID"`
	CategoryID uint
}
