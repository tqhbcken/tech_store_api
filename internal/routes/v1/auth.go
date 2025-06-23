package v1

import (
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/logout", middlewares.JWTAuthMiddleware(jwt.NewJWTConfig()), handlers.Logout)
		auth.POST("/refresh", handlers.RefreshToken)
		//DEBUG 
		// auth.GET("/test-redis", handlers.TestRedis)
		// auth.DELETE("/clear-redis", handlers.ClearRedis)

		// email. password
		// auth.POST("/reset-password", handlers.ResetPassword)
		// auth.POST("/verify-email", handlers.VerifyEmail)
	}
}
