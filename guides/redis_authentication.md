# Redis Authentication Implementation

## Tổng quan

Hệ thống authentication với Redis được implement để cung cấp:
- Session management
- Token blacklisting
- Rate limiting
- Refresh token management
- User activity tracking

## Cấu trúc Redis Keys

### 1. User Sessions
```
session:user:{user_id}
```
- Lưu trữ thông tin session của user
- Expiration: Theo JWT token duration
- Data: JSON object chứa user info, login time, last activity, IP, User-Agent

### 2. Refresh Tokens
```
refresh_token:user:{user_id}
```
- Lưu trữ refresh token cho user
- Expiration: Theo JWT refresh token duration
- Data: Refresh token string

### 3. Blacklisted Tokens
```
blacklist:token:{token_hash}
```
- Lưu trữ các token đã bị blacklist
- Expiration: Theo thời gian còn lại của token
- Data: "blacklisted"

### 4. Login Attempts (Rate Limiting)
```
login_attempts:{ip_address}
```
- Đếm số lần login thất bại từ IP
- Expiration: 15 phút
- Data: Số lần thử

### 5. User Login Info
```
login_info:user:{user_id}
```
- Lưu trữ thông tin login của user
- Expiration: 24 giờ
- Data: JSON object chứa login time, IP, User-Agent

## API Endpoints

### Authentication
- `POST /api/v1/auth/login` - Đăng nhập
- `POST /api/v1/auth/register` - Đăng ký
- `POST /api/v1/auth/logout` - Đăng xuất
- `POST /api/v1/auth/refresh` - Refresh token
- `GET /api/v1/auth/profile` - Lấy thông tin profile
- `GET /api/v1/auth/validate` - Validate token

### Request/Response Examples

#### Login
```json
POST /api/v1/auth/login
{
  "email": "user@example.com",
  "password": "password123"
}

Response:
{
  "code": 200,
  "status": "success",
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 86400
  }
}
```

#### Refresh Token
```json
POST /api/v1/auth/refresh
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

Response:
{
  "code": 200,
  "status": "success",
  "message": "Token refreshed successfully",
  "data": {
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expires_in": 86400
  }
}
```

#### Logout
```json
POST /api/v1/auth/logout
Authorization: Bearer {access_token}
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

Response:
{
  "code": 200,
  "status": "success",
  "message": "Logout successful"
}
```

## Security Features

### 1. Rate Limiting
- Giới hạn 5 lần login thất bại trong 15 phút
- Reset counter khi login thành công
- Block IP khi vượt quá giới hạn

### 2. Token Blacklisting
- Blacklist access token và refresh token khi logout
- Kiểm tra blacklist trước khi validate token
- Tự động expire theo thời gian token

### 3. Session Management
- Lưu trữ session data trong Redis
- Update last activity khi sử dụng token
- Tự động expire session

### 4. Refresh Token Rotation
- Tạo refresh token mới mỗi lần refresh
- Invalidate refresh token cũ
- Secure token generation

## Environment Variables

```bash
# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here
JWT_DURATION=24h
JWT_REFRESH_DURATION=168h

# Rate Limiting
LOGIN_ATTEMPTS_LIMIT=5
LOGIN_ATTEMPTS_DURATION=15m
```

## Middleware Usage

### JWT Authentication Middleware
```go
// Protected routes
protected := r.Group("/api/v1")
protected.Use(middlewares.JWTAuthMiddleware(authService))
{
    // Routes requiring authentication
}
```

### Role-based Authorization
```go
// Admin only routes
admin := protected.Group("/admin")
admin.Use(middlewares.RequireRole("admin"))
{
    // Admin routes
}
```

## Error Handling

### Common Error Responses
```json
{
  "code": 401,
  "status": "error",
  "message": "Invalid token"
}

{
  "code": 401,
  "status": "error", 
  "message": "Token has been revoked"
}

{
  "code": 429,
  "status": "error",
  "message": "Too many login attempts. Please try again later."
}
```

## Performance Considerations

1. **Redis Connection Pooling**: Sử dụng connection pool để tối ưu performance
2. **Key Expiration**: Tự động expire keys để tránh memory leak
3. **Batch Operations**: Sử dụng pipeline cho multiple Redis operations
4. **Caching Strategy**: Cache user data để giảm database queries

## Monitoring

### Redis Metrics
- Memory usage
- Connection count
- Command latency
- Key count by type

### Application Metrics
- Login success/failure rate
- Token validation count
- Session count
- Rate limiting events

## Troubleshooting

### Common Issues
1. **Redis Connection Failed**: Kiểm tra Redis service và connection string
2. **Token Validation Failed**: Kiểm tra JWT secret và expiration
3. **Session Not Found**: Kiểm tra Redis data và expiration
4. **Rate Limiting Too Strict**: Điều chỉnh limits trong config

### Debug Commands
```bash
# Check Redis keys
redis-cli keys "session:user:*"
redis-cli keys "blacklist:token:*"

# Check specific user session
redis-cli get "session:user:1"

# Check login attempts
redis-cli get "login_attempts:192.168.1.1"
``` 