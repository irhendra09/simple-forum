package model

import "time"

type RefreshToken struct {
	ID        int64     `gorm:"primaryKey;column:id" json:"id"`
	Token     string    `gorm:"column:token" json:"token"`
	UserID    int64     `gorm:"column:user_id" json:"user_id"`
	ExpiresAt time.Time `gorm:"column:expires_at" json:"expires_at"`
	Revoked   bool      `gorm:"column:revoked" json:"revoked"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
