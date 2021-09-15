package app

import (
	"github.com/Gunnsteinn/cryptoGuild/controllers/adventurer"
	"github.com/Gunnsteinn/cryptoGuild/controllers/healthCheck"
	"github.com/Gunnsteinn/cryptoGuild/controllers/scheduling"
	"github.com/Gunnsteinn/cryptoGuild/controllers/sponsors"
)

// mapUrls is used to mapping urls
func mapUrls() {
	router.GET("/AppStatus", healthCheck.AppStatus)

	router.GET("/AdventurerStatus/:id", adventurer.GetInfo)
	router.POST("/AllAdventurerStatus", adventurer.GetAllInfo)

	router.GET("/sponsors/:wallet_id", sponsors.GetByWalletId)
	router.GET("/sponsors/filterBy/:nick_name", sponsors.GetByQuery)
	router.GET("/sponsors", sponsors.GetAll)
	router.POST("/sponsor", sponsors.Create)
	router.PUT("/sponsors/:sponsors_id", sponsors.Update)
	router.DELETE("/sponsors/:wallet_id", sponsors.Delete)

	router.GET("/CronJobStart", scheduling.CronJobStart)
	router.GET("/CronJobStop", scheduling.CronJobStop)
}
