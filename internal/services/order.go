package services

import (
	"api_techstore/internal/models"

	"gorm.io/gorm"
)

type OrderService interface {
	GetAllOrders() ([]models.Order, error)
	GetOrderByID(id string) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	UpdateOrder(id string, order models.Order) (models.Order, error)
	DeleteOrder(id string) error
	GetOrdersByUserID(userID string) ([]models.Order, error)
}

type orderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{db: db}
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := s.db.Preload("User").Preload("OrderItems").Preload("ShippingAddress").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *orderService) GetOrderByID(id string) (models.Order, error) {
	var order models.Order
	if err := s.db.Preload("User").Preload("OrderItems").Preload("ShippingAddress").First(&order, "id = ?", id).Error; err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func (s *orderService) CreateOrder(order models.Order) (models.Order, error) {
	if err := s.db.Create(&order).Error; err != nil {
		return models.Order{}, err
	}

	// Preload related data after creation
	if err := s.db.Preload("User").Preload("OrderItems").Preload("ShippingAddress").First(&order, order.ID).Error; err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (s *orderService) UpdateOrder(id string, order models.Order) (models.Order, error) {
	var existingOrder models.Order
	if err := s.db.First(&existingOrder, "id = ?", id).Error; err != nil {
		return models.Order{}, err
	}
	if err := s.db.Save(&order).Error; err != nil {
		return models.Order{}, err
	}

	// Preload related data after update
	if err := s.db.Preload("User").Preload("OrderItems").Preload("ShippingAddress").First(&order, id).Error; err != nil {
		return models.Order{}, err
	}

	return order, nil
}

func (s *orderService) DeleteOrder(id string) error {
	var order models.Order
	if err := s.db.First(&order, "id = ?", id).Error; err != nil {
		return err
	}
	if err := s.db.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}

func (s *orderService) GetOrdersByUserID(userID string) ([]models.Order, error) {
	var orders []models.Order
	if err := s.db.Preload("User").Preload("OrderItems").Preload("ShippingAddress").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}
