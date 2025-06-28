#!/bin/bash

echo "🚀 Setting up TechStore API project..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

echo "✅ Go is installed"

# Install dependencies
echo "📦 Installing Go dependencies..."
go mod download

# Install swag CLI if not installed
if ! command -v swag &> /dev/null; then
    echo "📦 Installing swag CLI..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate Swagger documentation
echo "📚 Generating Swagger documentation..."
swag init -g cmd/main.go -o docs

# Make scripts executable
echo "🔧 Making scripts executable..."
chmod +x scripts/*.sh

echo "✅ Setup completed successfully!"
echo ""
echo "📋 Next steps:"
echo "1. Copy .env.example to .env and configure your environment variables"
echo "2. Start the server: go run cmd/main.go"
echo "3. Access Swagger UI: http://localhost:8080/swagger/index.html"
echo ""
echo "🎉 Happy coding!" 