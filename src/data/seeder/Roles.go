package seeder

import (
	"foodshop/data/models"

	"gorm.io/gorm"
)

func MigrateRoles(db *gorm.DB) {
	var role models.UserRoles

	role.State = "admin"
	role.ID = 1
	db.Model(&models.UserRoles{}).Create(&role)

	role.State = "chef"
	role.ID = 2
	db.Model(&models.UserRoles{}).Create(&role)

	role.State = "user"
	role.ID = 3
	db.Model(&models.UserRoles{}).Create(&role)

}
