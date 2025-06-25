package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAddressRoutes(r *gin.RouterGroup, ctn *container.Container) {
	addresses := r.Group("/addresses")
	{
		addresses.POST("", func(ctx *gin.Context) {
			handlers.CreateAddress(ctx, ctn)
		})
		addresses.GET("", func(ctx *gin.Context) {
			handlers.GetAddresses(ctx, ctn)
		})
		addresses.GET("/:id", func(ctx *gin.Context) {
			handlers.GetAddressByID(ctx, ctn)
		})
		addresses.PUT("/:id", func(ctx *gin.Context) {
			handlers.UpdateAddress(ctx, ctn)
		})
		addresses.DELETE("/:id", func(ctx *gin.Context) {
			handlers.DeleteAddress(ctx, ctn)
		})
	}
}
