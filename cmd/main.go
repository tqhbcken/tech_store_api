package main

import (
	"api_techstore/internal/database"
	"api_techstore/internal/models"
	"api_techstore/internal/routes"
	"api_techstore/pkg/logger"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	logger.InitLogger()

	// Khởi tạo database
	dbConn, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Tự động migrate các bảng dựa trên model
	if err := dbConn.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Brand{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()
	routes.SetupRouter(r)

	//chay server
	r.Run(":8080")
}
