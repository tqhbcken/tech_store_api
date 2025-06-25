package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupBrandRoute(r *gin.RouterGroup, ctn *container.Container) {
	brands := r.Group("/brands")
	{
		brands.GET("", func(c *gin.Context) {
			handlers.GetAllBrands(c, ctn)
		})
		brands.GET("/:id", func(c *gin.Context) {
			handlers.GetBrandById(c, ctn)
		})
		brands.POST("", 
		middlewares.ValidateRequest(&models.BrandCreateRequest{}),
		func(c *gin.Context) {
			handlers.CreateBrand(c, ctn)
		})
		brands.PUT("/:id", 
		middlewares.ValidateRequest(&models.BrandUpdateRequest{}),
		func(c *gin.Context) {
			handlers.UpdateBrand(c, ctn)
		})
		brands.DELETE("/:id", func(c *gin.Context) {
			handlers.DeleteBrand(c, ctn)
		})

		// brands.GET("/:slug", handlers.GetBrandBySlug)
		// brands.GET("/:id/products", handlers.GetProductsByBrand)
	}
}
