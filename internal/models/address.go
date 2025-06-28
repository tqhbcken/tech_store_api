package models

///root struct
type Address struct {
	Base
	UserID       uint   `gorm:"column:user_id;not null" json:"user_id"`
	FullName     string `gorm:"column:full_name;not null" json:"full_name"`
	Phone        string `gorm:"column:phone;not null" json:"phone"`
	AddressLine1 string `gorm:"column:address_line1;not null" json:"address_line1"`
	AddressLine2 string `gorm:"column:address_line2" json:"address_line2"`
	City         string `gorm:"column:city;not null" json:"city"`
	District     string `gorm:"column:district" json:"district"`
	IsDefault    bool   `gorm:"column:is_default;default:false" json:"is_default"`

	// Relations
	User   User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	// Orders []Order `json:"orders,omitempty" gorm:"foreignKey:ShippingAddressID"`
}

type AddressCreateRequest struct {
	UserID       uint   `json:"user_id" binding:"required"`
	FullName     string `json:"full_name" binding:"required,min=2,max=100"`
	Phone        string `json:"phone" binding:"required,min=8,max=20"`
	AddressLine1 string `json:"address_line1" binding:"required,max=255"`
	AddressLine2 string `json:"address_line2" binding:"omitempty,max=255"`
	City         string `json:"city" binding:"required,max=100"`
	District     string `json:"district" binding:"omitempty,max=100"`
	IsDefault    *bool  `json:"is_default" binding:"omitempty"`
}

type AddressUpdateRequest struct {
	FullName     string `json:"full_name" binding:"omitempty,min=2,max=100"`
	Phone        string `json:"phone" binding:"omitempty,min=8,max=20"`
	AddressLine1 string `json:"address_line1" binding:"omitempty,max=255"`
	AddressLine2 string `json:"address_line2" binding:"omitempty,max=255"`
	City         string `json:"city" binding:"omitempty,max=100"`
	District     string `json:"district" binding:"omitempty,max=100"`
	IsDefault    *bool  `json:"is_default" binding:"omitempty"`
}
