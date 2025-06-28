package models

type ProductImage struct {
	Base
	ProductID uint   `gorm:"column:product_id;not null" json:"product_id"`
	ImageURL  string `gorm:"column:image_url;not null" json:"image_url"`
	IsMain    bool   `gorm:"column:is_main;default:false" json:"is_main"`
	SortOrder int    `gorm:"column:sort_order;default:0" json:"sort_order"`

	// Relations
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}
