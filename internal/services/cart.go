package services

import (
	"api_techstore/internal/models"

	"gorm.io/gorm"
)

type CartService interface {
	CreateCart(cart models.Cart) (models.Cart, error)
	GetCartByID(id uint) (models.Cart, error)
	GetCartByUserID(userID uint) (models.Cart, error)
	DeleteCart(id uint) error
}

type cartService struct {
	db *gorm.DB
}

func NewCartService(db *gorm.DB) CartService {
	return &cartService{db: db}
}

func (s *cartService) CreateCart(cart models.Cart) (models.Cart, error) {
	err := s.db.Create(&cart).Error
	return cart, err
}

func (s *cartService) GetCartByID(id uint) (models.Cart, error) {
	var cart models.Cart
	err := s.db.Preload("Items").First(&cart, id).Error
	return cart, err
}

func (s *cartService) GetCartByUserID(userID uint) (models.Cart, error) {
	var cart models.Cart
	err := s.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error
	return cart, err
}

func (s *cartService) DeleteCart(id uint) error {
	return s.db.Delete(&models.Cart{}, id).Error
}
