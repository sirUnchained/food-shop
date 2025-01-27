package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserName string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Phone    string `gorm:"not null;unique"`
}
