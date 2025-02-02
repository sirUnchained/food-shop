package models

import "time"

// "foodshop/data/postgres"

type Users struct {
	ID        uint    `gorm:"primaryKey"`
	UserName  string  `gorm:"not null;unique"`
	Password  string  `gorm:"not null"`
	Email     string  `gorm:"not null;unique"`
	Phone     string  `gorm:"not null;unique"`
	Roles     []Roles `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
