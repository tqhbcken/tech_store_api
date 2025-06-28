package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupProductRoute(r *gin.RouterGroup, ctn *container.Container) {
	products := r.Group("/products")
	{
		products.GET("", func(c *gin.Context) {
			handlers.GetAllProducts(c, ctn)
		})
		products.GET("/:id", func(c *gin.Context) {
			handlers.GetProductById(c, ctn)
		})
		products.POST("",
			middlewares.RequireRole("admin"),
			middlewares.ValidateRequest(&models.ProductCreateRequest{}),
			func(c *gin.Context) {
				handlers.CreateProduct(c, ctn)
			})
		products.PUT("/:id",
			middlewares.RequireRole("admin"),
			middlewares.ValidateRequest(&models.ProductUpdateRequest{}),
			func(c *gin.Context) {
				handlers.UpdateProduct(c, ctn)
			})
		products.DELETE("/:id",
			middlewares.RequireRole("admin"),
			func(c *gin.Context) {
				handlers.DeleteProduct(c, ctn)
			})

		// Nested routes for product images
		SetupProductImageRoutes(products, ctn)

		// products.GET("/:slug", handlers.GetProductBySlug) //get product with slug
		// products.GET("/search", handlers.SearchProducts) //search products
	}
}
