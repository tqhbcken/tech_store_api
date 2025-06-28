# Hướng dẫn test API Auth (Postman/cURL)

## 1. Đăng ký (Register)
- **POST** `/api/v1/auth/register`
```json
{
  "full_name": "Nguyen Van A",
  "email": "test@example.com",
  "phone": "0912345678",
  "password": "12345678",
  "role": "user",
  "is_active": true
}
```

## 2. Đăng nhập (Login)
- **POST** `/api/v1/auth/login`
```json
{
  "email": "test@example.com",
  "password": "12345678"
}
```

## 3. Refresh Token
- **POST** `/api/v1/auth/refresh`
```json
{
  "refresh_token": "<refresh_token từ login>"
}
```

## 4. Logout
- **POST** `/api/v1/auth/logout`
- **Headers:** Authorization: Bearer <access_token>
```json
{
  "refresh_token": "<refresh_token từ login>"
}
```

## 5. Test lỗi validate
- Đăng ký thiếu trường:
```json
{
  "email": "test@example.com"
}
```
- Kết quả: 400 Bad Request, message: "Validation failed", error: liệt kê các trường thiếu

---

**Bạn có thể import các request này vào Postman hoặc dùng cURL để kiểm thử.** 