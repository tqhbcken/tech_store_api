package services

import (
	"api_techstore/internal/models"

	"gorm.io/gorm"
)

type PaymentService interface {
	CreatePayment(payment models.Payment) (models.Payment, error)
	GetPaymentByOrderID(orderID uint) (models.Payment, error)
	UpdatePaymentStatus(orderID uint, status string) error
}

type paymentService struct {
	db *gorm.DB
}

func NewPaymentService(db *gorm.DB) PaymentService {
	return &paymentService{db: db}
}

func (s *paymentService) CreatePayment(payment models.Payment) (models.Payment, error) {
	err := s.db.Create(&payment).Error
	return payment, err
}

func (s *paymentService) GetPaymentByOrderID(orderID uint) (models.Payment, error) {
	var payment models.Payment
	err := s.db.Preload("Order").Where("order_id = ?", orderID).First(&payment).Error
	return payment, err
}

func (s *paymentService) UpdatePaymentStatus(orderID uint, status string) error {
	return s.db.Model(&models.Payment{}).Where("order_id = ?", orderID).Update("status", status).Error
}
