package models

type Product struct {
	Base
	Name        string  `gorm:"column:name" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Price       float64 `gorm:"column:price" json:"price"`
	Quantity    int     `gorm:"column:quantity" json:"quantity"`
	CategoryID  uint    `gorm:"column:category_id;not null" json:"category_id"`
	BrandID     *uint   `gorm:"column:brand_id" json:"brand_id,omitempty"`
	Slug        string  `gorm:"column:slug;unique" json:"slug"`
	IsActive    bool    `gorm:"column:is_active" json:"is_active"`

	// Relations
	Category   Category    `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	Brand      *Brand      `json:"brand,omitempty" gorm:"foreignKey:BrandID"`
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:ProductID"`
}

type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required,min=2,max=200"`
	Description string  `json:"description" binding:"omitempty,max=1000"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Quantity    int     `json:"quantity" binding:"required,gte=0"`
	CategoryID  uint    `json:"category_id" binding:"required"`
	BrandID     *uint   `json:"brand_id,omitempty" binding:"omitempty"`
	Slug        string  `json:"slug" binding:"required,min=2,max=100"`
	IsActive    *bool   `json:"is_active" binding:"omitempty"`
}

type ProductUpdateRequest struct {
	Name        string  `json:"name" binding:"omitempty,min=2,max=200"`
	Description string  `json:"description" binding:"omitempty,max=1000"`
	Price       float64 `json:"price" binding:"omitempty,gt=0"`
	Quantity    int     `json:"quantity" binding:"omitempty,gte=0"`
	CategoryID  uint    `json:"category_id" binding:"omitempty"`
	BrandID     *uint   `json:"brand_id,omitempty" binding:"omitempty"`
	Slug        string  `json:"slug" binding:"omitempty,min=2,max=100"`
	IsActive    *bool   `json:"is_active" binding:"omitempty"`
}
