package models

import "gorm.io/gorm"

type UserRoles struct {
	gorm.Model
	State string `gorm:"type:varchar(255)"`
	Users []User `gorm:"many2many:user_roles"`
}
