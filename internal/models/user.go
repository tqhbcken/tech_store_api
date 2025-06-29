package models

///root struct
type User struct {
	Base
	FullName     string `gorm:"column:full_name" json:"full_name"`
	Email        string `gorm:"column:email" json:"email"`
	Phone        string `gorm:"column:phone" json:"phone"`
	PasswordHash string `gorm:"column:password_hash;not null" json:"-"`
	Role         string `gorm:"column:role;default:user" json:"role"`
	IsActive     bool   `gorm:"column:is_active;default:false" json:"is_active"`

	// Relations
	Orders []Order `json:"orders,omitempty" gorm:"foreignKey:UserID"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type CreateUserReq struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=admin user" default:"user"`
	IsActive bool   `json:"is_active" binding:"omitempty" default:"false"`
}

type UserUpdateRequest struct {
	FullName string `json:"full_name" binding:"omitempty,min=2,max=100"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone" binding:"omitempty,min=8,max=20"`
	Password string `json:"password" binding:"omitempty,min=6"`
	Role     string `json:"role" binding:"omitempty,oneof=admin user"`
	IsActive *bool  `json:"is_active" binding:"omitempty"`
}
