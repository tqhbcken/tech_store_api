package unit

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/models"
	"api_techstore/test/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.ProductService)

	expectedProducts := []models.Product{
		{Name: "iPhone 15", Price: 1000},
		{Name: "MacBook Pro", Price: 2000},
	}
	mockService.On("GetAllProducts").Return(expectedProducts, nil)

	mockContainer := &container.Container{
		ProductService: mockService,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	handlers.GetAllProducts(c, mockContainer)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody struct {
		Message string           `json:"message"`
		Data    []models.Product `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Products retrieved successfully", responseBody.Message)
	assert.Equal(t, expectedProducts, responseBody.Data)

	mockService.AssertExpectations(t)
}
