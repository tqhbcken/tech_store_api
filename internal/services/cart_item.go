package services

import (
	"api_techstore/internal/models"

	"gorm.io/gorm"
)

type CartItemService interface {
	AddItemToCart(item models.CartItem) (models.CartItem, error)
	UpdateCartItem(id uint, item models.CartItem) (models.CartItem, error)
	RemoveItemFromCart(id uint) error
	GetItemsByCartID(cartID uint) ([]models.CartItem, error)
	ClearCart(cartID uint) error
}

type cartItemService struct {
	db *gorm.DB
}

func NewCartItemService(db *gorm.DB) CartItemService {
	return &cartItemService{db: db}
}

func (s *cartItemService) AddItemToCart(item models.CartItem) (models.CartItem, error) {
	// Kiểm tra xem item đã tồn tại trong cart chưa
	var existingItem models.CartItem
	err := s.db.Where("cart_id = ? AND product_id = ?", item.CartID, item.ProductID).First(&existingItem).Error

	if err == gorm.ErrRecordNotFound {
		// Item chưa tồn tại, tạo mới
		err = s.db.Create(&item).Error
		return item, err
	} else if err != nil {
		// Có lỗi khác
		return models.CartItem{}, err
	} else {
		// Item đã tồn tại, update quantity
		newQuantity := existingItem.Quantity + item.Quantity
		err = s.db.Model(&existingItem).Update("quantity", newQuantity).Error
		if err != nil {
			return models.CartItem{}, err
		}
		existingItem.Quantity = newQuantity
		return existingItem, nil
	}
}

func (s *cartItemService) UpdateCartItem(id uint, item models.CartItem) (models.CartItem, error) {
	var existing models.CartItem
	if err := s.db.First(&existing, id).Error; err != nil {
		return models.CartItem{}, err
	}
	if err := s.db.Model(&existing).Updates(item).Error; err != nil {
		return models.CartItem{}, err
	}
	return existing, nil
}

func (s *cartItemService) RemoveItemFromCart(id uint) error {
	return s.db.Delete(&models.CartItem{}, id).Error
}

func (s *cartItemService) GetItemsByCartID(cartID uint) ([]models.CartItem, error) {
	var items []models.CartItem
	err := s.db.Where("cart_id = ?", cartID).Find(&items).Error
	return items, err
}

func (s *cartItemService) ClearCart(cartID uint) error {
	return s.db.Where("cart_id = ?", cartID).Delete(&models.CartItem{}).Error
}
