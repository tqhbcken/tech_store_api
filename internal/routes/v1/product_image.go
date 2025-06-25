package v1

import (
	"api_techstore/internal/container"
	// "api_techstore/internal/handlers"
	// "api_techstore/internal/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupProductImageRoutes configures routes for product images, likely nested under products
func SetupProductImageRoutes(r *gin.RouterGroup, ctn *container.Container) {
	// These routes are nested under /products/:productId and should be protected
	// images := r.Group("/:productId/images", middlewares.RequireRole("admin"))
	// {
	// 	images.POST("", func(ctx *gin.Context) {
	// 		handlers.AddProductImage(ctx, ctn)
	// 	})
	// 	images.DELETE("/:imageId", func(ctx *gin.Context) {
	// 		handlers.DeleteProductImage(ctx, ctn)
	// 	})
	// }
}
