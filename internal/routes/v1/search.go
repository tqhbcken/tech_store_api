package v1

import (
	"api_techstore/internal/container"

	"github.com/gin-gonic/gin"
)

func SetupSearchRoute(r *gin.RouterGroup, ctn *container.Container) {
	search := r.Group("/search")
	{
		search.GET("", func(c *gin.Context) {
		})
		search.GET("/suggestions", func(c *gin.Context) {
		})
		search.GET("/filters", func(c *gin.Context) {
		})
		search.GET("/popular", func(c *gin.Context) {
		})
	}
}