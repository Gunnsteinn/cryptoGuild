package adventurer

import (
	"github.com/Gunnsteinn/cryptoGuild/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Info(c *gin.Context) {
	userID := c.Param("id")
	if len(userID) <= 0 {
		c.JSON(http.StatusBadRequest, "ID param is empty")
		return
	}

	result, getErr := services.GetAdventurerDetail(userID)
	if getErr != nil {
		c.JSON(http.StatusInternalServerError, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
