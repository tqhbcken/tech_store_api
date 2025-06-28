# Error Handling Guide

## Tổng quan

Project đã được cải thiện error handling với hệ thống structured errors, giúp:
- ✅ Không expose internal errors ra client
- ✅ Có error codes chuẩn
- ✅ Response format nhất quán
- ✅ Dễ debug và maintain

## Cấu trúc Error Response

### Response Format
```json
{
  "code": 400,
  "status": "error",
  "message": "Validation failed",
  "data": null,
  "error": {
    "code": "VALIDATION_FAILED",
    "message": "Validation failed",
    "details": "Email format is invalid",
    "context": {
      "validation_errors": [
        {
          "field": "email",
          "message": "Invalid email format"
        }
      ]
    }
  }
}
```

## Error Codes

### Authentication Errors
- `INVALID_CREDENTIALS` - Email/password không đúng
- `TOKEN_EXPIRED` - Token đã hết hạn
- `TOKEN_INVALID` - Token không hợp lệ
- `TOKEN_REVOKED` - Token đã bị thu hồi
- `UNAUTHORIZED` - Chưa xác thực
- `FORBIDDEN` - Không có quyền

### Validation Errors
- `VALIDATION_FAILED` - Validation thất bại
- `INVALID_INPUT` - Input không hợp lệ
- `MISSING_FIELD` - Thiếu field bắt buộc

### Resource Errors
- `NOT_FOUND` - Resource không tồn tại
- `ALREADY_EXISTS` - Resource đã tồn tại
- `RESOURCE_DELETED` - Resource đã bị xóa

### Database Errors
- `DATABASE_ERROR` - Lỗi database
- `CONNECTION_FAILED` - Lỗi kết nối
- `TRANSACTION_FAILED` - Lỗi transaction

### External Service Errors
- `REDIS_ERROR` - Lỗi Redis
- `PAYMENT_FAILED` - Lỗi thanh toán
- `EMAIL_SEND_FAILED` - Lỗi gửi email

## Cách sử dụng trong Handlers

### 1. Basic Error Handling
```go
// ❌ Cũ - Expose internal error
response.ErrorResponse(c, http.StatusInternalServerError, err.Error())

// ✅ Mới - Structured error
response.HandleError(c, err)
```

### 2. Specific Error Types
```go
// Database errors
response.DatabaseErrorResponse(c, err)

// Redis errors
response.RedisErrorResponse(c, err)

// Not found
response.NotFoundResponse(c, "Product")

// Validation errors
response.ValidationErrorResponse(c, "Email format is invalid")
```

### 3. Custom App Errors
```go
// Tạo custom error
appErr := apperrors.New(apperrors.ErrCodeInsufficientStock, "Insufficient stock", http.StatusBadRequest)
response.NewErrorResponse(c, appErr)

// Hoặc sử dụng predefined
response.NewErrorResponse(c, apperrors.NewInvalidCredentials())
response.NewErrorResponse(c, apperrors.NewAlreadyExists("Email"))
```

### 4. Error với Context
```go
appErr := apperrors.NewValidationFailed("Validation failed")
appErr.Context = map[string]interface{}{
    "validation_errors": validationErrors,
}
response.NewErrorResponse(c, appErr)
```

## Ví dụ thực tế

### Login Handler
```go
func Login(c *gin.Context, ctn *container.Container) {
    user, err := services.Login(ctn.DB, req.Email, req.Password)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            response.NewErrorResponse(c, apperrors.NewInvalidCredentials())
            return
        }
        response.HandleError(c, err)
        return
    }
    // ... rest of logic
}
```

### Product Handler
```go
func GetProductById(c *gin.Context, ctn *container.Container) {
    product, err := ctn.ProductService.GetProductById(id)
    if err != nil {
        if err == gorm.ErrRecordNotFound {
            response.NotFoundResponse(c, "Product")
            return
        }
        response.DatabaseErrorResponse(c, err)
        return
    }
    response.SuccessResponse(c, http.StatusOK, "Product retrieved successfully", product)
}
```

## Middleware Integration

### Validation Middleware
Tự động handle validation errors với structured response:
```go
// Tự động tạo validation error response
middlewares.ValidateRequest(&models.LoginReq{})
```

### JWT Middleware
Tự động handle authentication errors:
```go
// Tự động tạo auth error response
middlewares.JWTAuthMiddleware(jwtConfig)
```

## Best Practices

### 1. Không expose internal errors
```go
// ❌ Bad
response.ErrorResponse(c, http.StatusInternalServerError, err.Error())

// ✅ Good
response.HandleError(c, err)
```

### 2. Log errors chi tiết
```go
// Log error details cho debugging
ctn.Logger.WithError(err).Error("Failed to create user")
response.HandleError(c, err)
```

### 3. Sử dụng specific error types
```go
// ✅ Specific error handling
if err == gorm.ErrRecordNotFound {
    response.NotFoundResponse(c, "User")
    return
}
```

### 4. Consistent error messages
```go
// ✅ Consistent messages
response.NewErrorResponse(c, apperrors.NewInvalidCredentials())
response.NewErrorResponse(c, apperrors.NewAlreadyExists("Email"))
```

## Testing Error Responses

### Test Error Scenarios
```go
func TestLogin_InvalidCredentials(t *testing.T) {
    // Test invalid credentials
    // Should return INVALID_CREDENTIALS error
}

func TestGetProduct_NotFound(t *testing.T) {
    // Test product not found
    // Should return NOT_FOUND error
}
```

## Migration Guide

### Từ Error Handling Cũ
```go
// ❌ Cũ
if err != nil {
    response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
    return
}

// ✅ Mới
if err != nil {
    response.HandleError(c, err)
    return
}
```

### Từ Specific Error Handling
```go
// ❌ Cũ
if err == gorm.ErrRecordNotFound {
    response.ErrorResponse(c, http.StatusNotFound, "User not found")
    return
}

// ✅ Mới
if err == gorm.ErrRecordNotFound {
    response.NotFoundResponse(c, "User")
    return
}
```

## Kết luận

Hệ thống error handling mới cung cấp:
- **Consistency**: Format response nhất quán
- **Security**: Không expose internal errors
- **Maintainability**: Dễ debug và maintain
- **User Experience**: Error messages rõ ràng cho client
- **Developer Experience**: Error codes chuẩn và dễ sử dụng 