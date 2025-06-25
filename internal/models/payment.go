package models

type Payment struct {
	Base
	OrderID       uint    `json:"order_id" gorm:"not null;unique"`
	Amount        float64 `json:"amount" gorm:"not null"`
	Method        string  `json:"method" gorm:"not null"` // momo, zalopay, vnpay, cod
	Status        string  `json:"status" gorm:"default:pending"` // pending, completed, failed, refunded
	
	Order Order `json:"order,omitempty" gorm:"foreignKey:OrderID"`
}