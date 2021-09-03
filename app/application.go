package app

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	mapUrls()
	err := router.Run()
	if err != nil {
		return
	}
}
