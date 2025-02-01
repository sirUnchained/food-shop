package models

type Roles struct {
	ID     uint   `gorm:"primaryKey"`
	State  string `gorm:"size:50"`
	UserID uint
}
