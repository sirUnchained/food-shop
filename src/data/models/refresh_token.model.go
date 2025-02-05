package models

import (
	"time"

	"gorm.io/gorm"
)

type RefreshTokens struct {
	ID        uint  `gorm:"primaryKey"`
	Token     int64 `gorm:"unique"`
	UserID    uint
	User      Users
	CreatedAt time.Time `gorm:"index"`
	ExpiresAt time.Time `gorm:"index"`
}

// before create row sets the ExpiresAt field to 30 days from now
func (token *RefreshTokens) BeforeCreate(tx *gorm.DB) (err error) {
	token.ExpiresAt = time.Now().Add(30 * 24 * time.Hour)
	return
}
