package models

type Address struct {
	Base
	UserID       uint   `json:"user_id" gorm:"not null"`
	FullName     string `json:"full_name" gorm:"not null"`
	Phone        string `json:"phone" gorm:"not null"`
	AddressLine1 string `json:"address_line1" gorm:"not null"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city" gorm:"not null"`
	District     string `json:"district"`
	IsDefault    bool   `json:"is_default" gorm:"default:false"`
	
	// Relations
	User   User    `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Orders []Order `json:"orders,omitempty" gorm:"foreignKey:ShippingAddressID"`
}
