package v1

import (
	"github.com/gin-gonic/gin"
	"api_techstore/internal/handlers"
)

func SetupOrderRoute(route *gin.RouterGroup) {
	order := route.Group("/order")
	{
		order.GET("/", handlers.GetAllOrders)
		order.POST("/", handlers.CreateOrder)
		order.GET("/:id", handlers.GetOrderByID)
		order.PUT("/:id", handlers.UpdateOrder)
		order.DELETE("/:id", handlers.DeleteOrder)
		// order.GET("/user/:userId", handlers.GetOrdersByUserID)
		// order.GET("/status/:status", handlers.GetOrdersByStatus)
	}
}