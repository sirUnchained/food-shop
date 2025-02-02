package models

import "time"

type Restaurants struct {
	ID          uint `gorm:"primaryKey"`
	Description string
	Address     string
	PostalCode  string
	IsVerify    bool
	CategoryID  uint    `gorm:"index"`
	Foods       []Foods `gorm:"foreignKey:RestaurantID"`
	Owner       uint
	User        Users `gorm:"foreignKey:Owner"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
