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

// func (u *Users) BeforeDelete(tx *gorm.DB) (err error) {
// 	db := postgres.GetDb()

// 	userRole := new(Roles)
// 	db.Model(&Roles{}).Where("user_id = ?", u.ID).First(userRole)

// 	if userRole.State == "admin" {
// 		return errors.New("admin user not allowed to delete")
// 	}
// 	return
// }
