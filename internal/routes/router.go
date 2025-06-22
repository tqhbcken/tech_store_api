package routes

import (
	"api_techstore/internal/middlewares"
	v1 "api_techstore/internal/routes/v1"
	"api_techstore/pkg/jwt"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// Global middleware 
	r.Use(middlewares.RequestLogger())

	
	routeV1 := r.Group("/api/v1")
	{
		// Auth routes (không cần JWT)
		v1.SetupAuthRoute(routeV1)

		// Protected routes (cần JWT)
		protected := routeV1.Group("")
		protected.Use(middlewares.JWTAuthMiddleware(jwt.NewJWTConfig()))
		{
			v1.SetupUserRoute(protected)
			v1.SetupCategoryRoute(protected)
			v1.SetupBrandRoute(protected)
			v1.SetupProductRoute(protected)
			v1.SetupOrderRoute(protected)
		}
	}

	
	r.NoRoute(func(c *gin.Context) {
		response.ErrorResponse(c, http.StatusNotFound, "Page not found")
	})

	
	r.NoMethod(func(c *gin.Context) {
		response.ErrorResponse(c, http.StatusMethodNotAllowed, "Method not allowed")
	})

	//swagger
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
