package main

import (
	"api_techstore/internal/container"
	"api_techstore/internal/models"
	"api_techstore/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// init container
	ctn := container.NewContainer()
	
	// auto migrate
	if err := ctn.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Brand{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// init router
	r := gin.Default()
	routes.SetupRouter(r, ctn)

	// run server
	r.Run(":8080")
}
