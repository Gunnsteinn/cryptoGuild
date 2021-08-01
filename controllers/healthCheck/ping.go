package healthCheck

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AppStatus(c *gin.Context) {
	c.JSON(http.StatusOK,
		gin.H{
			"status": "OK",
		})
}
