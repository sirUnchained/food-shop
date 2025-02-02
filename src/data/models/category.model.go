package models

import "time"

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Slug        string `gorm:"not null;unique"`
	Title       string `gorm:"not null;unique"`
	CreatorID   uint   `gorm:"not null"`
	Creator     Users  `gorm:"foreignKey:CreatorID"`
	Restaurants []Restaurants
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
