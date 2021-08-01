package app

import "github.com/Gunnsteinn/cryptoGuild/controllers/healthCheck"

// mapUrls is used to mapping urls
func mapUrls() {
	router.GET("/AppStatus", healthCheck.AppStatus)
}
