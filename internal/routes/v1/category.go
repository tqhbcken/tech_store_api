package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupCategoryRoute(r *gin.RouterGroup, ctn *container.Container) {
	category := r.Group("/categories")
	{
		category.GET("", func(c *gin.Context) {
			handlers.GetAllCategories(c, ctn)
		})
		category.GET("/:id", func(c *gin.Context) {
			handlers.GetCategoryById(c, ctn)
		})
		category.POST("", 
		middlewares.ValidateRequest(&models.CategoryCreateRequest{}),
		func(c *gin.Context) {
			handlers.CreateCategory(c, ctn)
		})
		category.PUT("/:id", 
		middlewares.ValidateRequest(&models.CategoryUpdateRequest{}),
		func(c *gin.Context) {
			handlers.UpdateCategory(c, ctn)
		})
		category.DELETE("/:id", func(c *gin.Context) {
			handlers.DeleteCategory(c, ctn)
		})

		// category.GET("/:slug", handlers.GetCategoryBySlug)
	}
}
