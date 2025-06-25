package models

type ProductImage struct {
	Base
	ProductID uint   `json:"product_id" gorm:"not null"`
	ImageURL  string `json:"image_url" gorm:"not null"`
	IsMain    bool   `json:"is_main" gorm:"default:false"`
	SortOrder int    `json:"sort_order" gorm:"default:0"`
	
	// Relations
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

