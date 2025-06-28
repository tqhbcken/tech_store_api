package models

type Order struct {
	Base
	UserID      int     `gorm:"column:user_id" json:"user_id"`
	TotalAmount float64 `gorm:"column:total_amount" json:"total_amount"`
	Status      string  `gorm:"column:status" json:"status"`

	// Relations
	User       User        `json:"user,omitempty" gorm:"foreignKey:UserID"`
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID"`
}

type OrderCreateRequest struct {
	UserID      int     `json:"user_id" binding:"required,gt=0"`
	TotalAmount float64 `json:"total_amount" binding:"required,gt=0"`
	Status      string  `json:"status" binding:"required,oneof=pending completed cancelled"`
}

type OrderUpdateRequest struct {
	UserID      int     `json:"user_id" binding:"omitempty,gt=0"`
	TotalAmount float64 `json:"total_amount" binding:"omitempty,gt=0"`
	Status      string  `json:"status" binding:"omitempty,oneof=pending completed cancelled"`
}
