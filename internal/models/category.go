package models

type CategoryReq struct {
	Name string `json:"name" binding:"required"`
	Slug string `json:"slug" binding:"required"`
}

type Category struct {
	Base
	Name        string `gorm:"column:name;type:varchar(100);not null;unique"`
	Slug        string `gorm:"column:slug;type:varchar(100);not null;unique"`
}
