package models

type CartItem struct {
	Base
	CartID    uint
	ProductID uint
	Quantity  int `gorm:"not null;default:1"`
}
