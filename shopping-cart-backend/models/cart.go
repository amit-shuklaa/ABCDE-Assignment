package models

import (
	"time"
)

type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `gorm:"unique;not null" json:"user_id"` // Unique ensures one cart per user
	Name      string     `json:"name"`                           // e.g., "Shopping Cart"
	Status    string     `json:"status"`                         // e.g., "active", "converted"
	CreatedAt time.Time  `json:"created_at"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items"` // THIS LINE MUST BE PRESENT AND UNCOMMENTED!
}
