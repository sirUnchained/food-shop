package models

import (

	// "foodshop/data/postgres"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	UserName string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Phone    string `gorm:"not null;unique"`
	Roles    []Roles
}

// func (u *Users) BeforeDelete(tx *gorm.DB) (err error) {
// 	db := postgres.GetDb()

// 	userRole := new(Roles)
// 	db.Model(&Roles{}).Where("user_id = ?", u.ID).First(userRole)

// 	if userRole.State == "admin" {
// 		return errors.New("admin user not allowed to delete")
// 	}
// 	return
// }
