package models

type Product struct {
	Base
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"description" gorm:"column:description"`
	Price       float64 `json:"price" gorm:"column:price"`
	Quantity    int     `json:"quantity" gorm:"column:quantity"`
	CategoryID  uint    `json:"category_id" gorm:"not null"`
	BrandID     *uint   `json:"brand_id,omitempty"`
	Slug        string  `json:"slug" gorm:"unique"`
	IsActive    bool    `json:"is_active" gorm:"column:is_active"`

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
