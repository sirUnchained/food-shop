package models

import "gorm.io/gorm"

type Restaurants struct {
	gorm.Model
	Description string
	Address     string
	PostalCode  string
	IsVerify    bool
	CategoryID  uint    `gorm:"index"`
	Foods       []Foods `gorm:"foreignKey:RestaurantID"`
	Owner       uint
	User        Users `gorm:"foreignKey:Owner"`
}
