FROM golang:1.24.2-alpine

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]