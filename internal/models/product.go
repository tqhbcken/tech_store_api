package models

type Product struct {
	Base
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"description" gorm:"column:description"`
	Price       float64 `json:"price" gorm:"column:price"`
	Quantity    int     `json:"quantity" gorm:"column:quantity"`
	Category    string  `json:"category" gorm:"column:category"`
	// Image       string  `json:"image" gorm:"column:image"`
	IsActive    bool    `json:"is_active" gorm:"column:is_active"`
}