package mocks

import (
	"api_techstore/internal/models"

	"github.com/stretchr/testify/mock"
)

// BrandService is a mock type for the BrandService type
type BrandService struct {
	mock.Mock
}

// GetAllBrands provides a mock function
func (_m *BrandService) GetAllBrands() ([]models.Brand, error) {
	ret := _m.Called()
	return ret.Get(0).([]models.Brand), ret.Error(1)
}

// GetBrandById provides a mock function
func (_m *BrandService) GetBrandById(id string) (models.Brand, error) {
	ret := _m.Called(id)
	return ret.Get(0).(models.Brand), ret.Error(1)
}

// CreateBrand provides a mock function
func (_m *BrandService) CreateBrand(brand models.Brand) (models.Brand, error) {
	ret := _m.Called(brand)
	return ret.Get(0).(models.Brand), ret.Error(1)
}

// UpdateBrand provides a mock function
func (_m *BrandService) UpdateBrand(id string, brand models.Brand) (models.Brand, error) {
	ret := _m.Called(id, brand)
	return ret.Get(0).(models.Brand), ret.Error(1)
}

// DeleteBrand provides a mock function
func (_m *BrandService) DeleteBrand(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
