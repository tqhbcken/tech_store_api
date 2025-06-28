package services

import (
	"api_techstore/internal/models"

	"gorm.io/gorm"
)

type ProductService interface {
	GetAllProducts() ([]models.Product, error)
	GetProductById(id string) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(id string, product models.Product) (models.Product, error)
	DeleteProduct(id string) error
}

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{db: db}
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := s.db.Preload("Category").Preload("Brand").Find(&products).Error
	return products, err
}

func (s *productService) GetProductById(id string) (models.Product, error) {
	var product models.Product
	err := s.db.Preload("Category").Preload("Brand").First(&product, "id = ?", id).Error
	return product, err
}

func (s *productService) CreateProduct(product models.Product) (models.Product, error) {
	err := s.db.Create(&product).Error
	return product, err
}

func (s *productService) UpdateProduct(id string, product models.Product) (models.Product, error) {
	if err := s.db.Model(&models.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return models.Product{}, err
	}
	var updatedProduct models.Product
	if err := s.db.Preload("Category").Preload("Brand").First(&updatedProduct, "id = ?", id).Error; err != nil {
		return models.Product{}, err
	}
	return updatedProduct, nil
}

func (s *productService) DeleteProduct(id string) error {
	return s.db.Delete(&models.Product{}, "id = ?", id).Error
}
