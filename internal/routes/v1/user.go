package v1

import (
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	// "api_techstore/internal/middlewares"
	// "api_techstore/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func SetupUserRoute(r *gin.RouterGroup) {
	users := r.Group("/users")
	{
		users.GET("", middlewares.RequireRole("admin"), handlers.GetAllUsers)
		users.GET("/:id", middlewares.RequireRole("admin"), handlers.GetUserById)
		users.POST("", middlewares.RequireRole("admin"), handlers.CreateUser)
		users.PUT("/:id", middlewares.RequireRole("admin"), handlers.UpdateUser)
		users.DELETE("/:id", middlewares.RequireRole("admin"), handlers.DeleteUser)




		users.GET("/profile", handlers.GetUserProfile) 
		// users.POST("/:id/verify-email", handlers.VerifyEmail) 
	}


}