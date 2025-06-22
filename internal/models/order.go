package models

type CreateOrderReq struct {
	UserID      int     `json:"user_id" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
	Status      string  `json:"status" binding:"required,oneof=pending completed cancelled"`
}

type UpdateOrderReq struct {
	UserID      int     `json:"user_id" binding:"omitempty"`
	TotalAmount float64 `json:"total_amount" binding:"omitempty"`	
	Status      string  `json:"status" binding:"omitempty,oneof=pending completed cancelled"`
}


type Order struct {
	Base
	OrderID     int       `json:"order_id" column:"order_id"`
	UserID      int       `json:"user_id" column:"user_id"`
	TotalAmount float64   `json:"total_amount" column:"total_amount"`
	Status      string    `json:"status" column:"status"`
}