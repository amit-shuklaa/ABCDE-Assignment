package models

type CartItem struct {
	CartID uint `gorm:"primaryKey" json:"cart_id"`
	ItemID uint `gorm:"primaryKey" json:"item_id"`
}
