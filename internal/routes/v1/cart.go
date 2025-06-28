package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(r *gin.RouterGroup, ctn *container.Container) {
	cart := r.Group("/cart")
	{
		cart.GET("", middlewares.RequireRole("user", "admin"), func(ctx *gin.Context) {
			handlers.GetCart(ctx, ctn)
		})
		cart.POST("/items",
			middlewares.RequireRole("user", "admin"),
			middlewares.ValidateRequest(&models.CartAddItemRequest{}),
			func(ctx *gin.Context) {
				handlers.AddItemToCart(ctx, ctn)
			})
		cart.PUT("/items/:itemId",
			middlewares.RequireRole("user", "admin"),
			middlewares.ValidateRequest(&models.CartUpdateItemRequest{}),
			func(ctx *gin.Context) {
				handlers.UpdateCartItem(ctx, ctn)
			})
		cart.DELETE("/items/:itemId", middlewares.RequireRole("user", "admin"), func(ctx *gin.Context) {
			handlers.RemoveItemFromCart(ctx, ctn)
		})
		cart.DELETE("", 
		middlewares.RequireRole("user", "admin"), 
		func(ctx *gin.Context) {
			handlers.ClearCart(ctx, ctn)
		})
	}
}
