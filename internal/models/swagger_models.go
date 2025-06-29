package models

import "time"

// SwaggerBase represents the base model for Swagger documentation
// @Description Base model for Swagger documentation
type SwaggerBase struct {
	ID        uint      `json:"id" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`
}

// SwaggerUser represents user model for Swagger documentation
// @Description User model for Swagger documentation
type SwaggerUser struct {
	SwaggerBase
	FullName string `json:"full_name" example:"John Doe"`
	Email    string `json:"email" example:"john@example.com"`
	Phone    string `json:"phone" example:"0912345678"`
	Role     string `json:"role" example:"user"`
	IsActive bool   `json:"is_active" example:"true"`
}

// SwaggerProduct represents product model for Swagger documentation
// @Description Product model for Swagger documentation
type SwaggerProduct struct {
	SwaggerBase
	Name        string  `json:"name" example:"iPhone 15"`
	Description string  `json:"description" example:"Latest iPhone model"`
	Price       float64 `json:"price" example:"999.99"`
	Quantity    int     `json:"quantity" example:"10"`
	CategoryID  uint    `json:"category_id" example:"1"`
	BrandID     *uint   `json:"brand_id,omitempty" example:"1"`
	Slug        string  `json:"slug" example:"iphone-15"`
	IsActive    bool    `json:"is_active" example:"true"`
}

// SwaggerCategory represents category model for Swagger documentation
// @Description Category model for Swagger documentation
type SwaggerCategory struct {
	SwaggerBase
	Name string `json:"name" example:"Laptop"`
	Slug string `json:"slug" example:"laptop"`
}

// SwaggerBrand represents brand model for Swagger documentation
// @Description Brand model for Swagger documentation
type SwaggerBrand struct {
	SwaggerBase
	Name        string `json:"name" example:"Apple"`
	Description string `json:"description" example:"Apple Inc."`
	IsActive    bool   `json:"is_active" example:"true"`
	Slug        string `json:"slug" example:"apple"`
}

// SwaggerOrder represents order model for Swagger documentation
// @Description Order model for Swagger documentation
type SwaggerOrder struct {
	SwaggerBase
	UserID            uint    `json:"user_id" example:"1"`
	TotalAmount       float64 `json:"total_amount" example:"1999.99"`
	Status            string  `json:"status" example:"pending"` // pending, confirmed, processing, shipped, delivered, cancelled
	ShippingAddressID *uint   `json:"shipping_address_id,omitempty" example:"1"`
}

// SwaggerAddress represents address model for Swagger documentation
// @Description Address model for Swagger documentation
type SwaggerAddress struct {
	SwaggerBase
	UserID       uint   `json:"user_id" example:"1"`
	FullName     string `json:"full_name" example:"John Doe"`
	Phone        string `json:"phone" example:"0912345678"`
	AddressLine1 string `json:"address_line1" example:"123 Main St"`
	AddressLine2 string `json:"address_line2" example:"Apt 4B"`
	City         string `json:"city" example:"Hanoi"`
	District     string `json:"district" example:"Ba Dinh"`
	IsDefault    bool   `json:"is_default" example:"false"`
}

// SwaggerCart represents cart model for Swagger documentation
// @Description Cart model for Swagger documentation
type SwaggerCart struct {
	SwaggerBase
	UserID    *uint             `json:"user_id,omitempty" example:"1"`
	SessionID string            `json:"session_id" example:"session123"`
	Status    string            `json:"status" example:"active"`
	Items     []SwaggerCartItem `json:"items,omitempty"`
}

// SwaggerCartItem represents cart item model for Swagger documentation
// @Description Cart item model for Swagger documentation
type SwaggerCartItem struct {
	SwaggerBase
	CartID    uint `json:"cart_id" example:"1"`
	ProductID uint `json:"product_id" example:"1"`
	Quantity  int  `json:"quantity" example:"2"`
}

// SwaggerPayment represents payment model for Swagger documentation
// @Description Payment model for Swagger documentation
type SwaggerPayment struct {
	SwaggerBase
	OrderID uint    `json:"order_id" example:"1"`
	Amount  float64 `json:"amount" example:"1999.99"`
	Method  string  `json:"method" example:"cod"`     // momo, zalopay, vnpay, cod
	Status  string  `json:"status" example:"pending"` // pending, completed, failed, refunded, cancelled
}
