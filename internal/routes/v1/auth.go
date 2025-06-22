package v1

import (
	"api_techstore/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.GET("/logout", handlers.Logout)
		auth.POST("/refresh", handlers.RefreshToken)

		// email. password
		// auth.POST("/reset-password", handlers.ResetPassword)
		// auth.POST("/verify-email", handlers.VerifyEmail)
	}
}
