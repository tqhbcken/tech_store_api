package models

type OrderItem struct {
	Base
	OrderID   uint    `gorm:"column:order_id;not null" json:"order_id"`
	ProductID uint    `gorm:"column:product_id;not null" json:"product_id"`
	Quantity  int     `gorm:"column:quantity;not null" json:"quantity"`
	UnitPrice float64 `gorm:"column:unit_price;type:numeric(10,2);not null" json:"unit_price"`

	// Relations
	Order   Order   `json:"order,omitempty" gorm:"foreignKey:OrderID"`
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
