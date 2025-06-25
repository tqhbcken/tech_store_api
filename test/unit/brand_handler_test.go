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

func TestGetAllBrands_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.BrandService)

	// Define expected data
	expectedBrands := []models.Brand{
		{Name: "Apple", Slug: "apple"},
		{Name: "Samsung", Slug: "samsung"},
	}

	// Mock the service call
	mockService.On("GetAllBrands").Return(expectedBrands, nil)

	// Create a mock container
	mockContainer := &container.Container{
		BrandService: mockService,
	}

	// Create a response recorder and context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the handler
	handlers.GetAllBrands(c, mockContainer)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var responseBody struct {
		Message string         `json:"message"`
		Data    []models.Brand `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Brands retrieved successfully", responseBody.Message)
	assert.Equal(t, expectedBrands, responseBody.Data)

	// Verify that the mock was called
	mockService.AssertExpectations(t)
}
