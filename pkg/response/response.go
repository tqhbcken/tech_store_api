package response

import (
	"github.com/gin-gonic/gin"
)

// Response represents the standard API response structure
// @Description Standard API response structure
type Response struct {
	Code    int         `json:"code" example:"200"`
	Status  string      `json:"status" example:"success"` // "success" | "error"
	Message string      `json:"message,omitempty" example:"Operation successful"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func SuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   nil,
		Status:  "success",
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Data:    nil,
		Error:   nil,
		Status:  "error",
	})
}
