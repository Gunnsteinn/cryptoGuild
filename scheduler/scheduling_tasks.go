package scheduler

import (
	"github.com/Gunnsteinn/cryptoGuild/services"
	"github.com/Gunnsteinn/cryptoGuild/utils/errors"
	"github.com/robfig/cron/v3"
	"log"
)

var job = cron.New()

func StartTask() *errors.RestErr {
	log.Println("Create new cron")
	//cron := cron.New()
	_, err := job.AddFunc("*/5 * * * *", func() {
		if errGetAndUpdate := getAndUpdate(); errGetAndUpdate != nil {
			log.Println(errGetAndUpdate)
		}
		log.Println("Cron executed successfully")
	})
	if err != nil {
		// Stop the scheduler (does not stop any jobs already running).
		job.Stop()
		return errors.NewInternalServerError(err.Error())
	}

	// Start cron with one scheduled job
	log.Println("Start cron")
	job.Start()

	return nil
}

func StopTask() *errors.RestErr {

	// Stop the scheduler (does not stop any jobs already running).
	job.Stop()

	// Start cron with one scheduled job
	log.Println("Stop cron")

	return nil
}

// GetAndUpdate all users.
func getAndUpdate() *errors.RestErr {
	user, getErr := services.SponsorsService.GetAllSponsor()
	if getErr != nil {
		return getErr
	}

	for _, sponsorModel := range user {
		_, saveErr := services.SponsorsService.UpdateSponsor(sponsorModel)
		if saveErr != nil {
			return saveErr
		}
	}
	return nil
}
