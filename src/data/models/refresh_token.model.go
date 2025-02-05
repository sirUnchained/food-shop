package models

import "time"

type RefresTokens struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      Users
	CreatedAt time.Time
}
