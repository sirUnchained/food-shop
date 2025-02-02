package models

import "time"

type Foods struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Price        string
	Description  string
	Pic          string
	IsAvailable  bool
	RestaurantID uint `gorm:"index"`
	UpdatedAt    time.Time
	CreatedAt    time.Time
}
