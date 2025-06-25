package unit

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetUserProfile_Success(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// Create a new response recorder
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a new request
	req, _ := http.NewRequest(http.MethodGet, "/users/profile", nil)
	c.Request = req

	// Mock the container and its dependencies if necessary
	// For now, let's assume a simple case and pass a nil container.
	// We might need to adjust this later.
	var ctn *container.Container = nil

	// Call the handler
	handlers.GetUserProfile(c, ctn)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	// We'll add more assertions about the body later.
}
 