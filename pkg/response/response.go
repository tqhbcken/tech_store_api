package response

import (
	"api_techstore/pkg/errors"

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

// AppErrorResponse represents error response structure
type AppErrorResponse struct {
	Code    string      `json:"code" example:"INVALID_CREDENTIALS"`
	Message string      `json:"message" example:"Invalid email or password"`
	Details string      `json:"details,omitempty" example:"Email format is invalid"`
	Context interface{} `json:"context,omitempty"`
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

// NewErrorResponse creates a structured error response
func NewErrorResponse(c *gin.Context, appErr *errors.AppError) {
	errorResp := AppErrorResponse{
		Code:    string(appErr.Code),
		Message: appErr.Message,
		Details: appErr.Details,
		Context: appErr.Context,
	}

	c.JSON(appErr.HTTPStatus, Response{
		Code:    appErr.HTTPStatus,
		Status:  "error",
		Message: appErr.Message,
		Data:    nil,
		Error:   errorResp,
	})
}

// HandleError handles any error and returns appropriate response
func HandleError(c *gin.Context, err error) {
	if appErr := errors.GetAppError(err); appErr != nil {
		NewErrorResponse(c, appErr)
		return
	}

	// Handle unknown errors
	NewErrorResponse(c, errors.NewInternalError(err))
}

// ValidationErrorResponse creates validation error response
func ValidationErrorResponse(c *gin.Context, details string) {
	appErr := errors.NewValidationFailed(details)
	NewErrorResponse(c, appErr)
}

// NotFoundResponse creates not found error response
func NotFoundResponse(c *gin.Context, resource string) {
	appErr := errors.NewNotFound(resource)
	NewErrorResponse(c, appErr)
}

// UnauthorizedResponse creates unauthorized error response
func UnauthorizedResponse(c *gin.Context) {
	NewErrorResponse(c, errors.NewUnauthorized())
}

// ForbiddenResponse creates forbidden error response
func ForbiddenResponse(c *gin.Context) {
	NewErrorResponse(c, errors.NewForbidden())
}

// DatabaseErrorResponse creates database error response
func DatabaseErrorResponse(c *gin.Context, err error) {
	appErr := errors.NewDatabaseError(err)
	NewErrorResponse(c, appErr)
}

// RedisErrorResponse creates Redis error response
func RedisErrorResponse(c *gin.Context, err error) {
	appErr := errors.NewRedisError(err)
	NewErrorResponse(c, appErr)
}
