package scheduling

import (
	"github.com/Gunnsteinn/cryptoGuild/scheduler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CronJobStart(c *gin.Context) {
	err := scheduler.Pepe()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "Run"})
}
