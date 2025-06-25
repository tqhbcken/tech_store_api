package routes

import (
	"api_techstore/internal/container"
	"api_techstore/internal/middlewares"
	v1 "api_techstore/internal/routes/v1"
	"api_techstore/pkg/jwt"
	"api_techstore/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, ctn *container.Container) {

	r.Use(middlewares.RequestLogger())

	routeV1 := r.Group("/api/v1")
	{
		// Auth routes (không cần JWT)
		v1.SetupAuthRoute(routeV1, ctn)

		// v1.SetupSearchRoute(routeV1, ctn)

		// Protected routes (cần JWT)
		protected := routeV1.Group("")
		protected.Use(middlewares.JWTAuthMiddleware(jwt.NewJWTConfig()))
		{
			v1.SetupUserRoute(protected, ctn)
			v1.SetupCategoryRoute(protected, ctn)
			v1.SetupBrandRoute(protected, ctn)
			v1.SetupProductRoute(protected, ctn)
			v1.SetupOrderRoute(protected, ctn)
			v1.SetupAddressRoutes(protected, ctn)
			v1.SetupCartRoutes(protected, ctn)
		}

		// Routes for both protected and public access
		v1.SetupPaymentRoutes(protected, routeV1, ctn)
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
