# Sử dụng Go 1.24.2 làm base image
FROM golang:1.24.2-alpine

# Cài đặt git và ca-certificates
RUN apk add --no-cache git ca-certificates

# Tạo thư mục làm việc
WORKDIR /app

# Copy file go.mod và go.sum trước
COPY go.mod go.sum ./

# Tải dependencies
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build ứng dụng
RUN go build -o main cmd/main.go

# Expose port 8080
EXPOSE 8080

# Chạy ứng dụng
CMD ["./main"] 