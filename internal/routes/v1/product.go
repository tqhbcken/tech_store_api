package v1

import (
	"github.com/gin-gonic/gin"
	"api_techstore/internal/handlers"
)

func SetupProductRoute(r *gin.RouterGroup) {
	products := r.Group("/products")
	{
		products.GET("", handlers.GetAllProducts)
		products.GET("/:id", handlers.GetProductById)
		products.POST("", handlers.CreateProduct)
		products.PUT("/:id", handlers.UpdateProduct)
		products.DELETE("/:id", handlers.DeleteProduct)
	}
}