package models

type CartItem struct {
	Base
	CartID    uint `gorm:"column:cart_id" json:"cart_id"`
	ProductID uint `gorm:"column:product_id" json:"product_id"`
	Quantity  int  `gorm:"column:quantity;not null;default:1" json:"quantity"`

	// Relations
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
