# Swagger Integration Guide

## Tổng quan

Project TechStore API đã được tích hợp Swagger để tạo tài liệu API tự động. Swagger UI cung cấp giao diện tương tác để test các API endpoints.

## Cài đặt

### 1. Cài đặt Swag CLI
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 2. Generate Swagger Documentation
```bash
# Chạy script tự động
./scripts/generate-swagger.sh

# Hoặc chạy lệnh trực tiếp
swag init -g cmd/main.go -o docs
```

## Sử dụng

### 1. Khởi động server
```bash
go run cmd/main.go
```

### 2. Truy cập Swagger UI
Mở trình duyệt và truy cập: `http://localhost:8080/swagger/index.html`

### 3. Test API
- Chọn endpoint muốn test
- Click "Try it out"
- Nhập thông tin cần thiết
- Click "Execute"

## Cấu trúc Annotations

### 1. API Information (main.go)
```go
// @title           TechStore API
// @version         1.0
// @description     A TechStore e-commerce API built with Go
// @host            localhost:8080
// @BasePath        /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

### 2. Handler Annotations
```go
// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.LoginReq true "Login credentials"
// @Success 200 {object} response.Response{data=map[string]interface{}} "Login successful"
// @Failure 400 {object} response.Response "Invalid request"
// @Router /auth/login [post]
func Login(c *gin.Context, ctn *container.Container) {
    // Handler implementation
}
```

## Tags được sử dụng

- **auth**: Authentication endpoints (login, register, logout, refresh)
- **users**: User management endpoints
- **products**: Product management endpoints
- **categories**: Category management endpoints
- **brands**: Brand management endpoints
- **orders**: Order management endpoints
- **cart**: Shopping cart endpoints
- **addresses**: Address management endpoints
- **payments**: Payment endpoints

## Security

Tất cả các protected endpoints đều sử dụng Bearer token authentication:
- Click "Authorize" button trên Swagger UI
- Nhập token theo format: `Bearer <your_jwt_token>`
- Token sẽ được áp dụng cho tất cả các requests

## Response Models

### Success Response
```json
{
  "code": 200,
  "status": "success",
  "message": "Operation successful",
  "data": {
    // Response data
  }
}
```

### Error Response
```json
{
  "code": 400,
  "status": "error",
  "message": "Error description",
  "error": [
    {
      "field": "email",
      "message": "Invalid email format"
    }
  ]
}
```

## Cập nhật Documentation

Khi thêm/sửa API endpoints:

1. **Thêm annotations** vào handler function
2. **Generate lại docs**:
   ```bash
   ./scripts/generate-swagger.sh
   ```
3. **Restart server** để áp dụng thay đổi

## Troubleshooting

### 1. Swagger UI không load
- Kiểm tra server đã chạy chưa
- Kiểm tra route `/swagger/*any` đã được setup
- Kiểm tra import `_ "api_techstore/docs"` trong main.go

### 2. Annotations không hiển thị
- Kiểm tra syntax annotations
- Generate lại docs: `swag init -g cmd/main.go -o docs`
- Restart server

### 3. Models không được recognize
- Đảm bảo models có struct tags đúng
- Kiểm tra import paths
- Generate lại docs

## Best Practices

1. **Luôn thêm annotations** cho mọi endpoint mới
2. **Sử dụng tags** để nhóm các endpoints liên quan
3. **Mô tả rõ ràng** parameters và responses
4. **Test API** thông qua Swagger UI trước khi deploy
5. **Cập nhật docs** mỗi khi thay đổi API

## Links hữu ích

- [Swaggo Documentation](https://github.com/swaggo/swag)
- [Swagger UI](https://swagger.io/tools/swagger-ui/)
- [OpenAPI Specification](https://swagger.io/specification/) 