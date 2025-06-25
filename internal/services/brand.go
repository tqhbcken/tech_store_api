package services

import (
	"api_techstore/internal/models"

	"gorm.io/gorm"
)

type BrandService interface {
	GetAllBrands() ([]models.Brand, error)
	GetBrandById(id string) (models.Brand, error)
	CreateBrand(brand models.Brand) (models.Brand, error)
	UpdateBrand(id string, brand models.Brand) (models.Brand, error)
	DeleteBrand(id string) error
}

type brandService struct {
	db *gorm.DB
}

func NewBrandService(db *gorm.DB) BrandService {
	return &brandService{db: db}
}

func (s *brandService) GetAllBrands() ([]models.Brand, error) {
	var brands []models.Brand
	err := s.db.Find(&brands).Error
	return brands, err
}

func (s *brandService) GetBrandById(id string) (models.Brand, error) {
	var brand models.Brand
	err := s.db.First(&brand, "id = ?", id).Error
	return brand, err
}

func (s *brandService) CreateBrand(brand models.Brand) (models.Brand, error) {
	err := s.db.Create(&brand).Error
	return brand, err
}

func (s *brandService) UpdateBrand(id string, brand models.Brand) (models.Brand, error) {
	if err := s.db.Model(&models.Brand{}).Where("id = ?", id).Updates(brand).Error; err != nil {
		return models.Brand{}, err
	}
	var updatedBrand models.Brand
	if err := s.db.First(&updatedBrand, "id = ?", id).Error; err != nil {
		return models.Brand{}, err
	}
	return updatedBrand, nil
}

func (s *brandService) DeleteBrand(id string) error {
	err := s.db.Delete(&models.Brand{}, "id = ?", id).Error
	return err
}
