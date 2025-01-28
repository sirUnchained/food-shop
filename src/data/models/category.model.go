package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Slug      string `gorm:"not null;unique"`
	Title     string `gorm:"not null;unique"`
	CreatorID uint   `gorm:"not null"`
	Creator   Users  `gorm:"foreignKey:CreatorID"`
}
