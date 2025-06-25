package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"

	// "api_techstore/internal/middlewares"
	// "api_techstore/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func SetupUserRoute(r *gin.RouterGroup, ctn *container.Container) {
	users := r.Group("/users")
	{
		users.GET("", middlewares.RequireRole("admin"), func(ctx *gin.Context) {
			handlers.GetAllUsers(ctx, ctn)
		})
		users.GET("/:id", middlewares.RequireRole("admin"), func(ctx *gin.Context) {
			handlers.GetUserById(ctx, ctn)
		})
		users.POST("", middlewares.RequireRole("admin"), func(ctx *gin.Context) {
			handlers.CreateUser(ctx, ctn)
		})
		users.PUT("/:id", middlewares.RequireRole("admin"), func(ctx *gin.Context) {
			handlers.UpdateUser(ctx, ctn)
		})
		users.DELETE("/:id", middlewares.RequireRole("admin"), func(ctx *gin.Context) {
			handlers.DeleteUser(ctx, ctn)
		})

		users.GET("/profile", func(ctx *gin.Context) {
			handlers.GetUserProfile(ctx, ctn)
		})
		
		// users.POST("/:id/verify-email", handlers.VerifyEmail)
	}


}