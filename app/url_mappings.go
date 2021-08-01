package app

import (
	"github.com/Gunnsteinn/cryptoGuild/controllers/adventurer"
	"github.com/Gunnsteinn/cryptoGuild/controllers/healthCheck"
)

// mapUrls is used to mapping urls
func mapUrls() {
	router.GET("/AppStatus", healthCheck.AppStatus)
	router.GET("/AdventurerStatus/:id", adventurer.GetInfo)

	router.POST("/AllAdventurerStatus", adventurer.GetAllInfo)
}
