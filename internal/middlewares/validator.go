package middlewares

import (
	"io"
	"net/http"

	// "regexp"

	"api_techstore/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidationError represents a validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidateRequest validates request body against a struct
func ValidateRequest(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(model); err != nil {
			handleError(c, err)
			return
		}

		if err := validate.Struct(model); err != nil {
			handleError(c, err)
			return
		}

		c.Set("validated_model", model)
		c.Next()
	}
}

// ValidateQuery validates query parameters against a struct
func ValidateQuery(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindQuery(model); err != nil {
			handleError(c, err)
			return
		}

		if err := validate.Struct(model); err != nil {
			handleError(c, err)
			return
		}

		c.Set("validated_query", model)
		c.Next()
	}
}

// ValidateForm validates form data against a struct
func ValidateForm(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBind(model); err != nil {
			handleError(c, err)
			return
		}

		if err := validate.Struct(model); err != nil {
			handleError(c, err)
			return
		}

		c.Set("validated_form", model)
		c.Next()
	}
}

// ValidateParams validates URL parameters against a struct
func ValidateParams(model interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindUri(model); err != nil {
			handleError(c, err)
			return
		}

		if err := validate.Struct(model); err != nil {
			handleError(c, err)
			return
		}

		c.Set("validated_params", model)
		c.Next()
	}
}

// handleError processes validation errors and returns a structured response
func handleError(c *gin.Context, err error) {
	var errors []ValidationError

	// Custom message for EOF (empty body)
	if err == io.EOF {
		errors = append(errors, ValidationError{
			Field:   "request",
			Message: "Request body is required",
		})
	} else if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   e.Field(),
				Message: getMessage(e.Field(), e.Tag()),
			})
		}
	} else {
		errors = append(errors, ValidationError{
			Field:   "request",
			Message: err.Error(),
		})
	}

	resp := response.Response{
		Code:    http.StatusBadRequest,
		Status:  "error",
		Message: "Validation failed",
		Data:    nil,
		Error:   errors,
	}
	c.JSON(http.StatusBadRequest, resp)
	c.Abort()
}

// getMessage returns a user-friendly validation message
func getMessage(field, tag string) string {
	switch tag {
	case "required":
		return field + " is required"
	case "email":
		return "Invalid email format"
	case "min":
		return field + " is too short"
	case "max":
		return field + " is too long"
	case "numeric":
		return field + " must be numeric"
	case "oneof":
		return field + " has invalid value"
	case "gte":
		return field + " must be greater than or equal to 0"
	case "gt":
		return field + " must be greater than 0"
	default:
		return "Invalid " + field
	}
}

// GetValidatedModel retrieves the validated model from context
func GetValidatedModel(c *gin.Context) interface{} {
	return c.MustGet("validated_model")
}

// GetValidatedQuery retrieves the validated query from context
func GetValidatedQuery(c *gin.Context) interface{} {
	if query, exists := c.Get("validated_query"); exists {
		return query
	}
	return nil
}

// GetValidatedForm retrieves the validated form from context
func GetValidatedForm(c *gin.Context) interface{} {
	if form, exists := c.Get("validated_form"); exists {
		return form
	}
	return nil
}

// GetValidatedParams retrieves the validated params from context
func GetValidatedParams(c *gin.Context) interface{} {
	if params, exists := c.Get("validated_params"); exists {
		return params
	}
	return nil
}

// RegisterCustomValidations registers custom validation functions
// func RegisterCustomValidations() {
// 	// Example: Custom validation for Vietnamese phone numbers
// 	validate.RegisterValidation("vietnamese_phone", validateVietnamesePhone)

// 	// Example: Custom validation for Vietnamese ID card
// 	validate.RegisterValidation("vietnamese_id", validateVietnameseID)
// }

// validateVietnamesePhone validates Vietnamese phone number format
// func validateVietnamesePhone(fl validator.FieldLevel) bool {
// 	phone := fl.Field().String()
// 	// Vietnamese phone number patterns: 03xx, 05xx, 07xx, 08xx, 09xx
// 	patterns := []string{
// 		"^03[0-9]{8}$", // 03xxxxxxxx
// 		"^05[0-9]{8}$", // 05xxxxxxxx
// 		"^07[0-9]{8}$", // 07xxxxxxxx
// 		"^08[0-9]{8}$", // 08xxxxxxxx
// 		"^09[0-9]{8}$", // 09xxxxxxxx
// 	}

// 	for _, pattern := range patterns {
// 		if matched, _ := regexp.MatchString(pattern, phone); matched {
// 			return true
// 		}
// 	}
// 	return false
// }

// validateVietnameseID validates Vietnamese ID card format
// func validateVietnameseID(fl validator.FieldLevel) bool {
// 	id := fl.Field().String()
// 	// Vietnamese ID card: 12 digits
// 	pattern := "^[0-9]{12}$"
// 	matched, _ := regexp.MatchString(pattern, id)
// 	return matched
// }
