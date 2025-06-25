package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(r *gin.RouterGroup, ctn *container.Container) {
	cart := r.Group("/cart")
	{
		cart.GET("", func(ctx *gin.Context) {
			handlers.GetCart(ctx, ctn)
		})
		cart.POST("/items", func(ctx *gin.Context) {
			handlers.AddItemToCart(ctx, ctn)
		})
		cart.PUT("/items/:itemId", func(ctx *gin.Context) {
			handlers.UpdateCartItem(ctx, ctn)
		})
		cart.DELETE("/items/:itemId", func(ctx *gin.Context) {
			handlers.RemoveItemFromCart(ctx, ctn)
		})
		cart.DELETE("", func(ctx *gin.Context) {
			handlers.ClearCart(ctx, ctn)
		})
	}
}
