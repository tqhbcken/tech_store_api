package models

type CreateBrandReq struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	IsActive	bool   `json:"is_active" binding:"omitempty" default:"false"`
	Slug		string `json:"slug" binding:"required"`
}

type UpdateBrandReq struct {
	Name        string `json:"name" binding:"omitempty"`
	Description string `json:"description"`
	IsActive	bool   `json:"is_active" binding:"omitempty" default:"false"`
	Slug		string `json:"slug" binding:"omitempty"`
}

type Brand struct {
	Base
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	// Image string
	IsActive    bool   `json:"is_active" gorm:"column:is_active" default:"false"`
	Slug        string `json:"slug" gorm:"column:slug"`
}