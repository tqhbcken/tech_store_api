#!/bin/bash

echo "🚀 Building TechStore API Docker image..."

# Build Docker image
docker build -t techstore-api .

echo "✅ Build completed!"
echo ""
echo "📋 To run the application:"
echo "   docker-compose up -d"
echo ""
echo "🌐 Access the API at: http://localhost:8080"
echo "📚 Swagger docs at: http://localhost:8080/swagger/index.html" 