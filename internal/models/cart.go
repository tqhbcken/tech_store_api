package models

type Cart struct {
	Base
	UserID    *uint      `gorm:"column:user_id" json:"user_id,omitempty"`
	SessionID string     `gorm:"column:session_id;index" json:"session_id"`
	Status    string     `gorm:"column:status;default:'active'" json:"status"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items,omitempty"`
}

type CartAddItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gte=1"`
}

type CartUpdateItemRequest struct {
	Quantity int `json:"quantity" binding:"required,gte=1"`
}
