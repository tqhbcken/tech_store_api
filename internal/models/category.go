package models

type Category struct {
	Base
	Name string `gorm:"column:name;type:varchar(100);not null;unique"`
	Slug string `gorm:"column:slug;type:varchar(100);not null;unique"`

	// Relations
	Products []Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}

type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
	Slug string `json:"slug" binding:"required,min=2,max=100"`
}

type CategoryUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=2,max=100"`
	Slug string `json:"slug" binding:"omitempty,min=2,max=100"`
}
