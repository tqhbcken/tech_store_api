package mocks

import (
	"api_techstore/internal/models"

	"github.com/stretchr/testify/mock"
)

type ProductService struct {
	mock.Mock
}

func (_m *ProductService) GetAllProducts() ([]models.Product, error) {
	ret := _m.Called()
	return ret.Get(0).([]models.Product), ret.Error(1)
}

func (_m *ProductService) GetProductById(id string) (models.Product, error) {
	ret := _m.Called(id)
	return ret.Get(0).(models.Product), ret.Error(1)
}

func (_m *ProductService) CreateProduct(product models.Product) (models.Product, error) {
	ret := _m.Called(product)
	return ret.Get(0).(models.Product), ret.Error(1)
}

func (_m *ProductService) UpdateProduct(id string, product models.Product) (models.Product, error) {
	ret := _m.Called(id, product)
	return ret.Get(0).(models.Product), ret.Error(1)
}

func (_m *ProductService) DeleteProduct(id string) error {
	ret := _m.Called(id)
	return ret.Error(0)
}
