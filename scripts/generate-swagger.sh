#!/bin/bash

# Generate Swagger documentation
echo "Generating Swagger documentation..."

# Check if swag is installed
if ! command -v swag &> /dev/null; then
    echo "swag is not installed. Installing..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate docs
swag init -g cmd/main.go -o docs

echo "Swagger documentation generated successfully!"
echo "You can now access the Swagger UI at: http://localhost:8080/swagger/index.html" 