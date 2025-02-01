package models

import (
	"gorm.io/gorm"
)

type Foods struct {
	gorm.Model
	Name         string
	Price        string
	Description  string
	Pic          string
	IsAvailable  bool
	RestaurantID uint `gorm:"index"`
}
