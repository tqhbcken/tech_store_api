package mocks

import (
	"api_techstore/internal/models"

	"github.com/stretchr/testify/mock"
)

// CategoryService is a mock type for the CategoryService type
type CategoryService struct {
	mock.Mock
}

// GetAllCategories provides a mock function
func (_m *CategoryService) GetAllCategories() ([]models.Category, error) {
	ret := _m.Called()
	return ret.Get(0).([]models.Category), ret.Error(1)
}

// GetCategoryById provides a mock function
func (_m *CategoryService) GetCategoryById(id string) (models.Category, error) {
	ret := _m.Called(id)
	return ret.Get(0).(models.Category), ret.Error(1)
}

// CreateCategory provides a mock function
func (_m *CategoryService) CreateCategory(category models.Category) (models.Category, error) {
	ret := _m.Called(category)
	return ret.Get(0).(models.Category), ret.Error(1)
}

// UpdateCategory provides a mock function
func (_m *CategoryService) UpdateCategory(id string, category models.Category) (models.Category, error) {
	ret := _m.Called(id, category)
	return ret.Get(0).(models.Category), ret.Error(1)
}

// DeleteCategory provides a mock function
func (_m *CategoryService) DeleteCategory(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
