# Validation Middleware - Hướng dẫn sử dụng

## Tổng quan

Validation middleware đơn giản để validate request body khi user gửi data lên server.

## Cách sử dụng

### Validate Request Body
```go
// Route
router.POST("/users", 
    middlewares.ValidateRequest(&middlewares.UserCreateRequest{}),
    handlers.CreateUser,
)

// Handler
func CreateUser(c *gin.Context) {
    request := middlewares.GetValidatedModel(c).(*middlewares.UserCreateRequest)
    // Process request...
}
```

### Sử dụng MiddlewareBuilder
```go
// Kết hợp với auth
middlewares := &middlewares.MiddlewareBuilder{}
middlewares = middlewares.
    WithAuth(jwtConfig).
    WithValidation(&middlewares.ProductCreateRequest{})

router.POST("/products", middlewares.Build()..., handlers.CreateProduct)
```

## Validation Tags

### Binding Tags (Gin)
- `required`: Field bắt buộc
- `omitempty`: Field có thể bỏ trống
- `email`: Validate email format
- `min=3`: Minimum length
- `max=50`: Maximum length
- `gt=0`: Greater than 0
- `gte=0`: Greater than or equal to 0

### Validate Tags (Validator)
- `required`: Field bắt buộc
- `omitempty`: Field có thể bỏ trống
- `email`: Validate email format
- `min=3`: Minimum length
- `max=50`: Maximum length
- `gt=0`: Greater than 0
- `gte=0`: Greater than or equal to 0

## Response Format

Khi validation fail:
```json
{
    "success": false,
    "message": "Validation failed",
    "errors": [
        {
            "field": "email",
            "message": "Invalid email format"
        },
        {
            "field": "password",
            "message": "password is too short"
        }
    ]
}
```

## Ví dụ thực tế

### User Registration
```go
// Route
router.POST("/auth/register", 
    middlewares.ValidateRequest(&middlewares.UserCreateRequest{}),
    handlers.Register,
)

// Handler
func Register(c *gin.Context) {
    request := middlewares.GetValidatedModel(c).(*middlewares.UserCreateRequest)
    
    user := &models.User{
        Username:  request.Username,
        Email:     request.Email,
        Password:  request.Password,
        FirstName: request.FirstName,
        LastName:  request.LastName,
    }
    
    // Save user...
}
```

### Product Creation
```go
// Route
router.POST("/products", 
    middlewares.ValidateRequest(&middlewares.ProductCreateRequest{}),
    handlers.CreateProduct,
)

// Handler
func CreateProduct(c *gin.Context) {
    request := middlewares.GetValidatedModel(c).(*middlewares.ProductCreateRequest)
    
    product := &models.Product{
        Name:        request.Name,
        Description: request.Description,
        Price:       request.Price,
        Stock:       request.Stock,
        CategoryID:  request.CategoryID,
        BrandID:     request.BrandID,
    }
    
    // Save product...
}
```

## Best Practices

1. **Luôn sử dụng cả binding và validate tags**
2. **Tạo struct riêng cho từng loại request**
3. **Sử dụng meaningful field names** 