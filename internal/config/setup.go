package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

func ServerInit() {
	r := gin.Default()

	LoadEnvVar()
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// run server
	r.Run(host + ":" + port)
}