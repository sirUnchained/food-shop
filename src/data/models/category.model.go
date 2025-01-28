package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Slug    string  `gorm:"not null;unique"`
	Title   string  `gorm:"not null;unique"`
	Creator []Users `gorm:"foreignKey:ID;not null"`
}
