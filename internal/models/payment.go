package models

type Payment struct {
	Base
	OrderID uint    `gorm:"column:order_id;not null;unique" json:"order_id"`
	Amount  float64 `gorm:"column:amount;not null" json:"amount"`
	Method  string  `gorm:"column:method;not null" json:"method"` // momo, zalopay, vnpay, cod
	Status  string  `gorm:"column:status;default:pending;check:status IN ('pending', 'completed', 'failed', 'refunded', 'cancelled')" json:"status"`

	Order Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
}

type PaymentCreateRequest struct {
	OrderID uint    `json:"order_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required,gt=0"`
	Method  string  `json:"method" binding:"required,oneof=momo zalopay vnpay cod"`
}

type PaymentUpdateRequest struct {
	Status string `json:"status" binding:"required,oneof=pending completed failed refunded cancelled"`
}
