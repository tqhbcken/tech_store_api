package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupAddressRoutes(r *gin.RouterGroup, ctn *container.Container) {
	addresses := r.Group("/addresses")
	{
		addresses.POST("",
			middlewares.RequireRole("user", "admin"),
			middlewares.ValidateRequest(&models.AddressCreateRequest{}),
			func(ctx *gin.Context) {
				handlers.CreateAddress(ctx, ctn)
			})
		addresses.GET("", middlewares.RequireRole("user", "admin"), func(ctx *gin.Context) {
			handlers.GetAddresses(ctx, ctn)
		})
		addresses.GET("/:id", middlewares.RequireRole("user", "admin"), func(ctx *gin.Context) {
			handlers.GetAddressByID(ctx, ctn)
		})
		addresses.PUT("/:id",
			middlewares.RequireRole("user", "admin"),
			middlewares.ValidateRequest(&models.AddressUpdateRequest{}),
			func(ctx *gin.Context) {
				handlers.UpdateAddress(ctx, ctn)
			})
		addresses.DELETE("/:id", middlewares.RequireRole("user", "admin"), func(ctx *gin.Context) {
			handlers.DeleteAddress(ctx, ctn)
		})
	}

	// Admin routes for managing all addresses
	adminAddresses := r.Group("/admin/addresses")
	adminAddresses.Use(middlewares.RequireRole("admin"))
	{
		adminAddresses.GET("", func(ctx *gin.Context) {
			handlers.GetAllAddressesAdmin(ctx, ctn)
		})
		adminAddresses.GET("/user/:userId", func(ctx *gin.Context) {
			handlers.GetAddressesByUserID(ctx, ctn)
		})
	}
}
