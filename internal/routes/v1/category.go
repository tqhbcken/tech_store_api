package v1

import (
	"api_techstore/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoute(r *gin.RouterGroup) {
	category := r.Group("/categories")
	{
		category.GET("", handlers.GetAllCategories)
		category.GET("/:id", handlers.GetCategoryById)
		category.POST("", handlers.CreateCategory)
		category.PUT("/:id", handlers.UpdateCategory)
		category.DELETE("/:id", handlers.DeleteCategory)
	}
}
