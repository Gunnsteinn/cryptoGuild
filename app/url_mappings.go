package app

import (
	"github.com/Gunnsteinn/cryptoGuild/controllers/adventurer"
	"github.com/Gunnsteinn/cryptoGuild/controllers/healthCheck"
	"github.com/Gunnsteinn/cryptoGuild/controllers/sponsors"
)

// mapUrls is used to mapping urls
func mapUrls() {
	router.GET("/AppStatus", healthCheck.AppStatus)
	router.GET("/AdventurerStatus/:id", adventurer.GetInfo)

	router.POST("/AllAdventurerStatus", adventurer.GetAllInfo)

	router.POST("/sponsors", sponsors.Create)
	router.GET("/sponsors/:wallet_id", sponsors.Get)
	router.PUT("/sponsors/:sponsors_id", sponsors.Update)
	router.PATCH("/sponsors/:sponsors_id", sponsors.Update)
	router.DELETE("/sponsors/:sponsors_id", sponsors.Delete)
	router.GET("/internal/sponsors/search", sponsors.Search)
}
