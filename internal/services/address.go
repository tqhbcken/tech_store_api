package services

import (
	"api_techstore/internal/models"

	"gorm.io/gorm"
)

type AddressService interface {
	GetAllAddresses(userID uint) ([]models.Address, error)
	GetAllAddressesAdmin() ([]models.Address, error)
	GetAddressByID(id uint) (models.Address, error)
	CreateAddress(address models.Address) (models.Address, error)
	UpdateAddress(id uint, address models.Address) (models.Address, error)
	DeleteAddress(id uint) error
}

type addressService struct {
	db *gorm.DB
}

func NewAddressService(db *gorm.DB) AddressService {
	return &addressService{db: db}
}

func (s *addressService) GetAllAddresses(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := s.db.Preload("User").Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (s *addressService) GetAllAddressesAdmin() ([]models.Address, error) {
	var addresses []models.Address
	err := s.db.Preload("User").Find(&addresses).Error
	return addresses, err
}

func (s *addressService) GetAddressByID(id uint) (models.Address, error) {
	var address models.Address
	err := s.db.Preload("User").First(&address, id).Error
	return address, err
}

func (s *addressService) CreateAddress(address models.Address) (models.Address, error) {
	err := s.db.Create(&address).Error
	if err != nil {
		return models.Address{}, err
	}

	err = s.db.Preload("User").First(&address, address.ID).Error
	return address, err
}

func (s *addressService) UpdateAddress(id uint, address models.Address) (models.Address, error) {
	var existing models.Address
	if err := s.db.First(&existing, id).Error; err != nil {
		return models.Address{}, err
	}
	if err := s.db.Model(&existing).Updates(address).Error; err != nil {
		return models.Address{}, err
	}
	return existing, nil
}

func (s *addressService) DeleteAddress(id uint) error {
	return s.db.Delete(&models.Address{}, id).Error
}
