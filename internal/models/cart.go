package models

type Cart struct {
	Base
	UserID    *uint
	SessionID string     `gorm:"index"`
	Status    string     `gorm:"default:'active'"`
	Items     []CartItem `gorm:"foreignKey:CartID"`
}
