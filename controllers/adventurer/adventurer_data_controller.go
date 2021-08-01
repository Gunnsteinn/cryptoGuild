package adventurer

import (
	"github.com/Gunnsteinn/cryptoGuild/domain"
	"github.com/Gunnsteinn/cryptoGuild/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetInfo(c *gin.Context) {
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

func GetAllInfo(c *gin.Context) {
	var body domain.BodyRequestListOfAdventurer
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body.")
		return
	}

	result, getAllErr := services.GetAllAdventurerDetail(body)
	if getAllErr != nil {
		c.JSON(http.StatusInternalServerError, getAllErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
