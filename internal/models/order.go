package models

type Order struct {
	Base
	UserID            uint    `gorm:"column:user_id" json:"user_id"`
	TotalAmount       float64 `gorm:"column:total_amount" json:"total_amount"`
	Status            string  `gorm:"column:status;check:status IN ('pending', 'confirmed', 'processing', 'shipped', 'delivered', 'cancelled')" json:"status"`
	ShippingAddressID *uint   `gorm:"column:shipping_address_id" json:"shipping_address_id"`

	// Relations
	User            User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderItems      []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
	ShippingAddress *Address    `json:"shipping_address,omitempty" gorm:"foreignKey:ShippingAddressID"`
}

type OrderCreateRequest struct {
	// UserID            uint    `json:"user_id" binding:"required,gt=0"`
	// TotalAmount       float64 `json:"total_amount" binding:"required,gt=0"`
	Status            string `json:"status" binding:"omitempty" default:"pending"`
	ShippingAddressID *uint  `json:"shipping_address_id" binding:"required"`
}

type OrderUpdateRequest struct {
	// UserID            uint    `json:"user_id" binding:"omitempty,gt=0"`           // Không cho phép thay đổi user_id
	// TotalAmount       float64 `json:"total_amount" binding:"omitempty,gt=0"`     // Không cho phép thay đổi total_amount
	Status            string `json:"status" binding:"omitempty,oneof=pending confirmed processing shipped delivered cancelled"`
	ShippingAddressID *uint  `json:"shipping_address_id" binding:"omitempty"`
}
