package services

import (
	"fmt"

	"github.com/robfig/cron"
	"github.com/santiagoh1997/weather-api/version-4/utils"
	"go.uber.org/zap"
)

// JobService interacts with the persistance layer
// and performs job related tasks
type JobService struct {
	Scheduler *cron.Cron
	logger    *zap.SugaredLogger
	ws        *WeatherService
}

// NewJobService returns a JobService
func NewJobService(ws *WeatherService, l *zap.SugaredLogger, c *cron.Cron) *JobService {
	return &JobService{
		c,
		l,
		ws,
	}
}

// NewJob schedules a new job to be performed hourly
func (js *JobService) NewJob(city, country string) *utils.APIError {
	if err := js.Scheduler.AddFunc("@hourly", func() {
		w, apiErr := js.ws.Get(city, country)
		if apiErr != nil {
			js.logger.Error(fmt.Sprintf("Scheduled job failed: %s", apiErr.Message))
			return
		}
		js.logger.Infof("Running scheduled job for %s", w.LocationName)
	}); err != nil {
		js.logger.Error(fmt.Sprintf("Failed to schedule new job: %s", err.Error()))
		return utils.NewInternalServerError()
	}
	return nil
}
