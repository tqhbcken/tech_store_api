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
2. **Cài Go modules:**
   ```bash
   go mod download
   ```
3. **Cấu hình biến môi trường:**
   - Copy `.env.example` thành `.env` và chỉnh sửa thông tin (DB, Redis, JWT...)
4. **Tạo database PostgreSQL:**
   - Có thể dùng Docker Compose: `docker-compose up -d`
   - DB sẽ tự động tạo qua migration SQL (xem thư mục `internal/database/migrations/`)
5. **Generate Swagger docs:**
   ```bash
   ./scripts/generate-swagger.sh
   ```
6. **Chạy ứng dụng:**
   ```bash
   go run cmd/main.go
   ```
   Hoặc dùng Docker Compose:
   ```bash
   docker-compose up --build
   ```

## Tài liệu API (Swagger)
- Truy cập: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
- Có thể test API trực tiếp, nhập Bearer token để thử các endpoint bảo vệ.
- Tài liệu luôn được generate từ code thực tế (xem hướng dẫn trong `guides/swagger_integration.md`).

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
├── go.mod, go.sum      # Quản lý dependency
└── README.md           # Tài liệu tổng quan
```

## Lưu ý bảo mật & vận hành
- **Không commit file .env thật lên repo.**
- JWT secret, DB password, Redis password phải đặt đủ mạnh khi deploy thật.
- Không expose port database/redis ra ngoài nếu không cần.
- Đảm bảo chạy migration trước khi chạy app lần đầu.
- Nếu chạy bằng Docker Compose, code chỉ update khi rebuild image.

## Testing
- Đã có unit test cho một số handler/service (thư mục `test/`).
- Có thể mở rộng test coverage cho service, middleware, error case.
- Hỗ trợ mock service cho test.

## Migration database
- Migration SQL nằm trong `internal/database/migrations/`.
- Chạy migration thủ công hoặc tích hợp tool ngoài (hiện chưa có script migrate tự động).

## Troubleshooting
- Nếu code thay đổi không áp dụng khi chạy Docker, hãy rebuild image: `docker-compose up --build`
- Nếu mất kết nối DB/Redis, kiểm tra lại biến môi trường và container.
- Swagger UI không hiển thị: kiểm tra đã generate docs và import `_ "api_techstore/docs"` trong main.go.
- Xem log lỗi chi tiết trong terminal hoặc log file (nếu cấu hình).

## Đóng góp & phát triển
- Mọi đóng góp, báo lỗi hoặc đề xuất vui lòng tạo issue hoặc pull request.
- Đọc kỹ các file hướng dẫn trong `guides/` để hiểu rõ về JWT, Redis, Validation, Error Handling, ...
- Đảm bảo tuân thủ best practice về bảo mật, error handling, validate khi thêm mới API.

## Tài liệu liên quan
- [Hướng dẫn JWT](guides/jwt_basic_implementation.md)
- [Hướng dẫn Redis Authentication](guides/redis_authentication.md)
- [Hướng dẫn Middleware Validation](guides/validation_middleware_guide.md)
- [Hướng dẫn Swagger Integration](guides/swagger_integration.md)
- [Error Handling Guide](guides/error_handling_guide.md)
- [Swagger API Spec](docs/swagger.yaml)

> Các tài liệu tự viết nằm trong thư mục `guides/`. Thư mục `docs/` chứa file swagger được generate tự động.

## Đóng góp
Mọi đóng góp, báo lỗi hoặc đề xuất vui lòng tạo issue hoặc pull request. 
