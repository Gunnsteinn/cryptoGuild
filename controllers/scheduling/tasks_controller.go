package scheduling

import (
	"github.com/Gunnsteinn/cryptoGuild/scheduler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CronJobStart(c *gin.Context) {
	err := scheduler.StartTask()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "Run"})
}

func CronJobStop(c *gin.Context) {
	err := scheduler.StopTask()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "Stop"})
}
