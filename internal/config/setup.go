package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func ServerInit(r *gin.Engine) {
	LoadEnvVar()
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	address := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Server starting on %s", address)

	// run server
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to start server on %s: %v", address, err)
	}
}
