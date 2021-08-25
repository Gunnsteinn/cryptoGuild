package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var router = gin.Default()

func StartApplication() {
	mapUrls()

	port := os.Getenv("port")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Defaulting to port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
