package models

import "time"

type Orders struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"not null;index"`
	RestaurantID uint   `gorm:"not null;index"`
	Address      string `gorm:"not null"`
	PostalCode   string `gorm:"not null"`
	IsDelivered  bool   `gorm:"default:false"`
	Stars        int    `gorm:"default:0"`
	User         Users
	Restaurant   Restaurants
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
