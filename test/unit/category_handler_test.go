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

func TestGetAllCategories_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	mockService := new(mocks.CategoryService)

	// Define expected data
	expectedCategories := []models.Category{
		{Name: "Laptops", Slug: "laptops"},
		{Name: "Smartphones", Slug: "smartphones"},
	}

	// Mock the service call
	mockService.On("GetAllCategories").Return(expectedCategories, nil)

	// Create a mock container
	mockContainer := &container.Container{
		CategoryService: mockService,
	}

	// Create a response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call the handler
	handlers.GetAllCategories(c, mockContainer)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	var responseBody struct {
		Message string            `json:"message"`
		Data    []models.Category `json:"data"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Categories retrieved successfully", responseBody.Message)
	assert.Equal(t, expectedCategories, responseBody.Data)

	// Verify that the mock was called
	mockService.AssertExpectations(t)
}
