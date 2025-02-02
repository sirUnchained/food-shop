package models

import "time"

type Orders struct {
	ID          uint        `gorm:"primaryKey"`
	User        uint        `gorm:"not null"`
	Users       Users       `gorm:"foreignKey:User"`
	Restaurant  uint        `gorm:"not null"`
	Restaurants Restaurants `gorm:"foreignKey:Restaurant"`
	Address     string      `gorm:"not null"`
	PostalCode  string      `gorm:"not null"`
	IsDelivered bool        `gorm:"default:false"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
