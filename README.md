# TechStore API

## Giới thiệu
Đây là project API cho hệ thống TechStore, được xây dựng bằng ngôn ngữ Go. API cung cấp các chức năng quản lý người dùng, sản phẩm, đơn hàng, giỏ hàng, thanh toán, v.v. phù hợp cho các hệ thống thương mại điện tử hiện đại.

## Tính năng chính
- Đăng ký, đăng nhập, xác thực JWT
- Quản lý người dùng, sản phẩm, thương hiệu, danh mục
- Quản lý đơn hàng, giỏ hàng, thanh toán
- Tích hợp Redis, RabbitMQ, Elk Stack(mở rộng)
- Middleware xác thực, rate limiting, logging, tracing
- Hỗ trợ migration database, unit test, mock service

## Hướng dẫn cài đặt nhanh
1. **Clone project:**
   ```bash
   git clone <repo-url>
   cd api_techstore
   ```
2. **Cài đặt Go modules:**
   ```bash
   go mod download
   ```
3. **Cấu hình biến môi trường:**
   - Copy file `.env.example` thành `.env` và chỉnh sửa thông tin cấu hình phù hợp.
4. **Chạy migration database:**
   ```bash
   ./scripts/migrate.sh
   ```
5. **Chạy ứng dụng:**
   ```bash
   go run cmd/main.go
   ```

## Cấu trúc thư mục
```
api_techstore/
├── cmd/                # Entry point của ứng dụng
├── internal/           # Business logic, handler, service, model, middleware
├── pkg/                # Package dùng chung (jwt, logger, response, ...)
├── docs/               # Chỉ chứa file swagger, tài liệu sinh tự động
├── guides/             # Chứa tài liệu tự viết (hướng dẫn, giải thích, ...)
├── scripts/            # Script hỗ trợ build, migrate, test
├── monitoring/         # Cấu hình monitoring (ELK, ...)
├── deployments/        # File triển khai (nếu có)
├── test/               # Unit test, mock, test utils
├── go.mod, go.sum      # Quản lý dependency
└── README.md           # Tài liệu tổng quan
```

## Tài liệu chi tiết
- [Hướng dẫn JWT](guides/jwt_basic_implementation.md)
- [Hướng dẫn Redis Authentication](guides/redis_authentication.md)
- [Hướng dẫn Middleware Validation](guides/validation_middleware_guide.md)
- [Swagger API Spec](docs/swagger.yaml)

> Các tài liệu tự viết nằm trong thư mục `guides/`. Thư mục `docs/` chỉ chứa file swagger hoặc tài liệu sinh tự động.

## Đóng góp
Mọi đóng góp, báo lỗi hoặc đề xuất vui lòng tạo issue hoặc pull request. 