package v1

import (
	"api_techstore/internal/container"
	"api_techstore/internal/handlers"
	"api_techstore/internal/middlewares"
	"api_techstore/internal/models"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoute(route *gin.RouterGroup, ctn *container.Container) {
	order := route.Group("/order")
	{
		order.GET("/", func(c *gin.Context) {
			handlers.GetAllOrders(c, ctn)
		})
		order.POST("/",
			middlewares.RequireRole("user", "admin"),
			middlewares.ValidateRequest(&models.OrderCreateRequest{}),
			func(c *gin.Context) {
				handlers.CreateOrder(c, ctn)
			})
		order.GET("/:id", func(c *gin.Context) {
			handlers.GetOrderByID(c, ctn)
		})
		order.PUT("/:id",
			middlewares.RequireRole("user", "admin"),
			middlewares.ValidateRequest(&models.OrderUpdateRequest{}),
			func(c *gin.Context) {
				handlers.UpdateOrder(c, ctn)
			})
		order.DELETE("/:id",
			middlewares.RequireRole("user", "admin"),
			func(c *gin.Context) {
				handlers.DeleteOrder(c, ctn)
			})
		order.GET("/user/:userId", func(c *gin.Context) {
			handlers.GetOrdersByUserID(c, ctn)
		})
		// order.GET("/status/:status", handlers.GetOrdersByStatus)
	}
}
