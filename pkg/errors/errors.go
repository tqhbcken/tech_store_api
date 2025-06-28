package errors

import (
	"fmt"
	"net/http"
)

// ErrorCode represents application error codes
type ErrorCode string



/// definition err
const (
	// Authentication errors
	ErrCodeInvalidCredentials ErrorCode = "INVALID_CREDENTIALS"
	ErrCodeTokenExpired       ErrorCode = "TOKEN_EXPIRED"
	ErrCodeTokenInvalid       ErrorCode = "TOKEN_INVALID"
	ErrCodeTokenRevoked       ErrorCode = "TOKEN_REVOKED"
	ErrCodeUnauthorized       ErrorCode = "UNAUTHORIZED"
	ErrCodeForbidden          ErrorCode = "FORBIDDEN"

	// Validation errors
	ErrCodeValidationFailed ErrorCode = "VALIDATION_FAILED"
	ErrCodeInvalidInput     ErrorCode = "INVALID_INPUT"
	ErrCodeMissingField     ErrorCode = "MISSING_FIELD"

	// Resource errors
	ErrCodeNotFound        ErrorCode = "NOT_FOUND"
	ErrCodeAlreadyExists   ErrorCode = "ALREADY_EXISTS"
	ErrCodeResourceDeleted ErrorCode = "RESOURCE_DELETED"

	// Database errors
	ErrCodeDatabaseError     ErrorCode = "DATABASE_ERROR"
	ErrCodeConnectionFailed  ErrorCode = "CONNECTION_FAILED"
	ErrCodeTransactionFailed ErrorCode = "TRANSACTION_FAILED"

	// Business logic errors
	ErrCodeInsufficientStock ErrorCode = "INSUFFICIENT_STOCK"
	ErrCodeInvalidQuantity   ErrorCode = "INVALID_QUANTITY"
	ErrCodeCartEmpty         ErrorCode = "CART_EMPTY"
	ErrCodeOrderCancelled    ErrorCode = "ORDER_CANCELLED"

	// External service errors
	ErrCodeRedisError      ErrorCode = "REDIS_ERROR"
	ErrCodePaymentFailed   ErrorCode = "PAYMENT_FAILED"
	ErrCodeEmailSendFailed ErrorCode = "EMAIL_SEND_FAILED"

	// Internal errors
	ErrCodeInternalError      ErrorCode = "INTERNAL_ERROR"
	ErrCodeServiceUnavailable ErrorCode = "SERVICE_UNAVAILABLE"
)

// AppError represents application-specific errors
type AppError struct {
	Code       ErrorCode              `json:"code"`
	Message    string                 `json:"message"`
	Details    string                 `json:"details,omitempty"`
	HTTPStatus int                    `json:"-"`
	Err        error                  `json:"-"`
	Context    map[string]interface{} `json:"context,omitempty"`
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%s)", e.Code, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap returns the underlying error
func (e *AppError) Unwrap() error {
	return e.Err
}

/// New creates a new AppError (contructor)
func New(code ErrorCode, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
	}
}

// NewWithError creates a new AppError with underlying error
func NewWithError(code ErrorCode, message string, httpStatus int, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
		Err:        err,
	}
}
///

// NewWithDetails creates a new AppError with additional details
func NewWithDetails(code ErrorCode, message, details string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		Details:    details,
		HTTPStatus: httpStatus,
	}
}

// Wrap wraps an existing error with AppError
func Wrap(err error, code ErrorCode, message string, httpStatus int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		HTTPStatus: httpStatus,
		Err:        err,
	}
}

// Common error constructors
func NewInvalidCredentials() *AppError {
	return New(ErrCodeInvalidCredentials, "Invalid email or password", http.StatusUnauthorized)
}

func NewTokenExpired() *AppError {
	return New(ErrCodeTokenExpired, "Token has expired", http.StatusUnauthorized)
}

func NewTokenInvalid() *AppError {
	return New(ErrCodeTokenInvalid, "Invalid token", http.StatusUnauthorized)
}

func NewTokenRevoked() *AppError {
	return New(ErrCodeTokenRevoked, "Token has been revoked", http.StatusUnauthorized)
}

func NewUnauthorized() *AppError {
	return New(ErrCodeUnauthorized, "Unauthorized access", http.StatusUnauthorized)
}

func NewForbidden() *AppError {
	return New(ErrCodeForbidden, "Insufficient permissions", http.StatusForbidden)
}

func NewNotFound(resource string) *AppError {
	return New(ErrCodeNotFound, fmt.Sprintf("%s not found", resource), http.StatusNotFound)
}

func NewAlreadyExists(resource string) *AppError {
	return New(ErrCodeAlreadyExists, fmt.Sprintf("%s already exists", resource), http.StatusConflict)
}

func NewValidationFailed(details string) *AppError {
	return NewWithDetails(ErrCodeValidationFailed, "Validation failed", details, http.StatusBadRequest)
}

func NewDatabaseError(err error) *AppError {
	return NewWithError(ErrCodeDatabaseError, "Database operation failed", http.StatusInternalServerError, err)
}

func NewRedisError(err error) *AppError {
	return NewWithError(ErrCodeRedisError, "Redis operation failed", http.StatusInternalServerError, err)
}

func NewInternalError(err error) *AppError {
	return NewWithError(ErrCodeInternalError, "Internal server error", http.StatusInternalServerError, err)
}

// IsAppError checks if an error is an AppError
func IsAppError(err error) bool {
	_, ok := err.(*AppError)
	return ok
}

// GetAppError extracts AppError from error chain
func GetAppError(err error) *AppError {
	if appErr, ok := err.(*AppError); ok {
		return appErr
	}
	return nil
}

// GetHTTPStatus returns HTTP status code for an error
func GetHTTPStatus(err error) int {
	if appErr := GetAppError(err); appErr != nil {
		return appErr.HTTPStatus
	}
	return http.StatusInternalServerError
}
