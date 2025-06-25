package models

type Brand struct {
	Base
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	IsActive    bool   `json:"is_active" gorm:"column:is_active" default:"false"`
	Slug        string `json:"slug" gorm:"column:slug"`

	// Relations
	Products []Product `json:"products,omitempty" gorm:"foreignKey:BrandID"`
}

type BrandCreateRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=255"`
	IsActive    *bool  `json:"is_active" binding:"omitempty"`
	Slug        string `json:"slug" binding:"required,min=2,max=100"`
}

type BrandUpdateRequest struct {
	Name        string `json:"name" binding:"omitempty,min=2,max=100"`
	Description string `json:"description" binding:"omitempty,max=255"`
	IsActive    *bool  `json:"is_active" binding:"omitempty"`
	Slug        string `json:"slug" binding:"omitempty,min=2,max=100"`
}
