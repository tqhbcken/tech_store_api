# JWT Basic Implementation

## Tổng quan

Hệ thống JWT cơ bản được implement với các tính năng:
- User authentication (login/register)
- JWT token generation và validation
- Protected routes với middleware
- Password hashing với bcrypt

## Cấu trúc API

### Authentication Endpoints

#### 1. Register User
```http
POST /api/v1/auth/register
Content-Type: application/json

{
  "full_name": "John Doe",
  "email": "john@example.com",
  "phone": "0123456789",
  "password": "password123",
  "role": "user",
  "is_active": true
}
```

**Response:**
```json
{
  "code": 201,
  "status": "success",
  "message": "User registered successfully"
}
```

#### 2. Login User
```http
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "code": 200,
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### 3. Logout User
```http
GET /api/v1/auth/logout
Authorization: Bearer {token}
```

**Response:**
```json
{
  "code": 200,
  "status": "success",
  "message": "Logout successful"
}
```

#### 4. Refresh Token
```http
POST /api/v1/auth/refresh
Authorization: Bearer {token}
```

**Response:**
```json
{
  "code": 200,
  "status": "success",
  "message": "Token refreshed successfully"
}
```

## Protected Routes

Tất cả các routes sau đây yêu cầu JWT token trong header:

### Users
- `GET /api/v1/users` - Lấy danh sách users
- `GET /api/v1/users/:id` - Lấy user theo ID
- `POST /api/v1/users` - Tạo user mới
- `PUT /api/v1/users/:id` - Cập nhật user
- `DELETE /api/v1/users/:id` - Xóa user

### Categories
- `GET /api/v1/categories` - Lấy danh sách categories
- `POST /api/v1/categories` - Tạo category mới
- `PUT /api/v1/categories/:id` - Cập nhật category
- `DELETE /api/v1/categories/:id` - Xóa category

### Brands
- `GET /api/v1/brands` - Lấy danh sách brands
- `POST /api/v1/brands` - Tạo brand mới
- `PUT /api/v1/brands/:id` - Cập nhật brand
- `DELETE /api/v1/brands/:id` - Xóa brand

### Products
- `GET /api/v1/products` - Lấy danh sách products
- `POST /api/v1/products` - Tạo product mới
- `PUT /api/v1/products/:id` - Cập nhật product
- `DELETE /api/v1/products/:id` - Xóa product

### Orders
- `GET /api/v1/orders` - Lấy danh sách orders
- `POST /api/v1/orders` - Tạo order mới
- `PUT /api/v1/orders/:id` - Cập nhật order
- `DELETE /api/v1/orders/:id` - Xóa order

## JWT Token Structure

### Header
```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

### Payload
```json
{
  "user_id": 1,
  "role": "user",
  "exp": 1640995200,
  "iat": 1640908800
}
```

### Claims
- `user_id`: ID của user
- `role`: Role của user (user/admin)
- `exp`: Expiration time
- `iat`: Issued at time

## Middleware Usage

### JWT Authentication Middleware
```go
// Áp dụng cho protected routes
protected := r.Group("/api/v1")
protected.Use(middlewares.JWTAuthMiddleware(jwt.NewJWTConfig()))
{
    // Protected routes here
}
```

### Role-based Authorization
```go
// Admin only routes
admin := protected.Group("/admin")
admin.Use(middlewares.RequireRole("admin"))
{
    // Admin routes here
}
```

## Environment Variables

```bash
# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here-change-this-in-production
JWT_DURATION=24h

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=techstore
DB_SSLMODE=disable
```

## Error Responses

### Authentication Errors
```json
{
  "code": 401,
  "status": "error",
  "message": "Authorization header is required"
}
```

```json
{
  "code": 401,
  "status": "error",
  "message": "Invalid authorization header format"
}
```

```json
{
  "code": 401,
  "status": "error",
  "message": "Invalid token"
}
```

```json
{
  "code": 401,
  "status": "error",
  "message": "Token has expired"
}
```

### Validation Errors
```json
{
  "code": 400,
  "status": "error",
  "message": "Email and password are required"
}
```

```json
{
  "code": 409,
  "status": "error",
  "message": "Email already exists"
}
```

## Security Features

1. **Password Hashing**: Sử dụng bcrypt để hash password
2. **JWT Secret**: Sử dụng secret key để sign tokens
3. **Token Expiration**: Tokens có thời gian hết hạn
4. **Protected Routes**: Middleware bảo vệ các routes nhạy cảm
5. **Role-based Access**: Phân quyền theo role

## Testing

### Test Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

### Test Protected Route
```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### Test Register
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "full_name": "Jane Doe",
    "email": "jane@example.com",
    "phone": "0987654321",
    "password": "password123",
    "role": "user",
    "is_active": true
  }'
```

## Troubleshooting

### Common Issues

1. **JWT Secret Not Set**: Đảm bảo JWT_SECRET được set trong environment
2. **Token Expired**: Token đã hết hạn, cần login lại
3. **Invalid Token Format**: Token không đúng format Bearer
4. **Database Connection**: Kiểm tra kết nối database

### Debug Commands

```bash
# Check JWT token
echo "YOUR_TOKEN" | base64 -d

# Test database connection
psql -h localhost -U postgres -d techstore

# Check server logs
tail -f logs/app.log
```

## Performance Considerations

1. **Token Size**: JWT tokens có thể lớn, cân nhắc sử dụng Redis cho session
2. **Database Queries**: Cache user data để giảm database queries
3. **Token Validation**: Validate token một lần và cache kết quả
4. **Password Hashing**: Sử dụng bcrypt với cost phù hợp

## Future Enhancements

1. **Refresh Tokens**: Implement refresh token mechanism
2. **Token Blacklisting**: Blacklist tokens khi logout
3. **Rate Limiting**: Giới hạn số lần login thất bại
4. **Multi-factor Authentication**: Thêm 2FA
5. **OAuth Integration**: Tích hợp với Google, Facebook, etc. 