package services

import (
	"api_techstore/internal/models"

	"gorm.io/gorm"
)

type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoryById(id string) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(id string, category models.Category) (models.Category, error)
	DeleteCategory(id string) error
}

type categoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return &categoryService{db: db}
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := s.db.Find(&categories).Error
	return categories, err
}

func (s *categoryService) GetCategoryById(id string) (models.Category, error) {
	var category models.Category
	err := s.db.First(&category, "id = ?", id).Error
	return category, err
}

func (s *categoryService) CreateCategory(category models.Category) (models.Category, error) {
	err := s.db.Create(&category).Error
	return category, err
}

func (s *categoryService) UpdateCategory(id string, category models.Category) (models.Category, error) {
	if err := s.db.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return models.Category{}, err
	}
	// Return the updated category
	var updatedCategory models.Category
	if err := s.db.First(&updatedCategory, "id = ?", id).Error; err != nil {
		return models.Category{}, err
	}
	return updatedCategory, nil
}

func (s *categoryService) DeleteCategory(id string) error {
	err := s.db.Delete(&models.Category{}, "id = ?", id).Error
	return err
}
