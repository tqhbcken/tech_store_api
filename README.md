# TechStore API

## Giới thiệu
TechStore API là backend RESTful cho hệ thống thương mại điện tử, xây dựng bằng Go (Gin, GORM), hỗ trợ quản lý người dùng, sản phẩm, đơn hàng, giỏ hàng, thanh toán, v.v. Dự án hướng tới kiến trúc rõ ràng, bảo mật, dễ mở rộng và dễ bảo trì.

## Tính năng nổi bật
- Đăng ký, đăng nhập, xác thực JWT, refresh token
- Quản lý người dùng, sản phẩm, thương hiệu, danh mục, đơn hàng, giỏ hàng, địa chỉ, thanh toán
- Middleware: xác thực, phân quyền, validate, logging, rate limiting (cơ bản)
- Error handling chuẩn hóa, không lộ lỗi nội bộ
- Tích hợp Redis (token/session), PostgreSQL, Docker Compose
- Tài liệu API tự động với Swagger UI
- Unit test cho handler/service (một phần)
- Hỗ trợ migration DB thủ công qua file SQL

## Cài đặt & Chạy nhanh

1. **Clone project:**
   ```bash
   git clone <repo-url>
   cd api_techstore
   ```

2. **Chạy với Docker Compose:**
   ```bash
   docker-compose up --build
   ```

3. **Truy cập:**
   - API: http://localhost:8082
   - Swagger UI: http://localhost:8082/swagger/index.html

## Cấu hình Docker Compose

Dự án sử dụng Docker Compose với 3 services:

- **postgres**: PostgreSQL 15 (port 5433)
- **redis**: Redis 7 (port 6380)  
- **api**: TechStore API (port 8082)

## Tài liệu API (Swagger)
- **URL**: http://localhost:8082/swagger/index.html
- Có thể test API trực tiếp, nhập Bearer token để thử các endpoint bảo vệ

## Cấu trúc thư mục
```
api_techstore/
├── cmd/                # Entry point
├── internal/           # Business logic: handler, service, model, middleware
├── pkg/                # Package dùng chung (jwt, logger, response, ...)
├── docs/               # Swagger docs (auto-generate)
├── guides/             # Tài liệu tự viết (JWT, Redis, Validation, ...)
├── scripts/            # Script build, swagger, setup
├── monitoring/         # Cấu hình monitoring (ELK, ...)
├── deployments/        # File triển khai (nếu có)
├── test/               # Unit test, mock, test utils
├── docker-compose.yml  # Docker Compose configuration
├── Dockerfile          # Docker image build
├── go.mod, go.sum      # Quản lý dependency
└── README.md           # Tài liệu tổng quan
```

## Lệnh Docker Compose hữu ích

```bash
# Chạy toàn bộ hệ thống
docker-compose up --build

# Chạy ở background
docker-compose up -d

# Xem logs
docker-compose logs -f api

# Dừng tất cả services
docker-compose down

# Rebuild image
docker-compose build --no-cache
```

## Lưu ý bảo mật & vận hành
- **Không commit file .env thật lên repo**
- JWT secret, DB password, Redis password phải đặt đủ mạnh khi deploy thật
- Database sẽ tự động migrate khi khởi động ứng dụng
- Code chỉ update khi rebuild image: `docker-compose up --build`

## Testing
- Đã có unit test cho một số handler/service (thư mục `test/`)
- Có thể mở rộng test coverage cho service, middleware, error case

## Troubleshooting

- **Code thay đổi không áp dụng**: Rebuild image: `docker-compose up --build`
- **Mất kết nối DB/Redis**: Kiểm tra healthcheck và restart services
- **Port conflict**: Kiểm tra port 8082, 5433, 6380 có đang được sử dụng không

## Đóng góp & phát triển
- Mọi đóng góp, báo lỗi hoặc đề xuất vui lòng tạo issue hoặc pull request
- Đọc kỹ các file hướng dẫn trong `guides/` để hiểu rõ về JWT, Redis, Validation, Error Handling, ...

## Tài liệu liên quan
- [Hướng dẫn JWT](guides/jwt_basic_implementation.md)
- [Hướng dẫn Redis Authentication](guides/redis_authentication.md)
- [Hướng dẫn Middleware Validation](guides/validation_middleware_guide.md)
- [Hướng dẫn Swagger Integration](guides/swagger_integration.md)
- [Error Handling Guide](guides/error_handling_guide.md)

## Đóng góp
Mọi đóng góp, báo lỗi hoặc đề xuất vui lòng tạo issue hoặc pull request.
