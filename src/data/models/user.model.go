package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	Admin Role = "admin"
	User  Role = "user"
	Chef  Role = "chef"
)

type Users struct {
	gorm.Model
	UserName string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Phone    string `gorm:"not null;unique"`
	// Roles    []string `gorm:"default:user"`
}
