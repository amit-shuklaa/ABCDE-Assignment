package models

import (
	"time"
)

type Order struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CartID    uint      `gorm:"unique;not null" json:"cart_id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
