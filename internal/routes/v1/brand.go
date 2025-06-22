package v1

import (
	"api_techstore/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupBrandRoute(r *gin.RouterGroup) {
	brands := r.Group("/brands")
	{
		brands.GET("", handlers.GetAllBrands)
		brands.GET("/:id", handlers.GetBrandById)
		brands.POST("", handlers.CreateBrand)
		brands.PUT("/:id", handlers.UpdateBrand)
		brands.DELETE("/:id", handlers.DeleteBrand)
	}
}