package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Here you can add your validation logic
		// For example, checking if required fields are present in the request body
		// or validating query parameters, headers, etc.

		

		// If validation fails, you can return an error response
		if !isValidRequest(c) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			c.Abort() // Stop further processing of the request
			return
		}

		c.Next() 
	}
}

func isValidRequest(c *gin.Context) bool {
	// Implement your validation logic here
	return true
}
