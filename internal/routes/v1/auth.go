package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(r *gin.RouterGroup, ctn *container.Container) {

	auth := r.Group("/auth")
	{
		auth.POST("/register",
			middlewares.ValidateRequest(&models.RegisterReq{}),
			func(ctx *gin.Context) {
				handlers.Register(ctx, ctn)
			})
		auth.POST("/login",
			middlewares.ValidateRequest(&models.LoginReq{}),
			func(ctx *gin.Context) {
				handlers.Login(ctx, ctn)
			})
		auth.POST("/logout",
			middlewares.JWTAuthMiddleware(ctn),
			func(ctx *gin.Context) {
				handlers.Logout(ctx, ctn)
			})
		auth.POST("/refresh", handlers.RefreshToken)

		//DEBUG
		auth.GET("/test-redis", handlers.TestRedis)
		auth.DELETE("/clear-redis", handlers.ClearRedis)

		// email. password //
		// auth.POST("/forgot-password", handlers.ForgotPassword)
		// auth.POST("/reset-password", handlers.ResetPassword)
		// auth.POST("/verify-email", handlers.VerifyEmail)
	}
}
