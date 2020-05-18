package services_test

import (
	"context"
	"testing"

	"github.com/robfig/cron"
	"github.com/santiagoh1997/weather-api/version-4/logger"
	"github.com/santiagoh1997/weather-api/version-4/services"
	"github.com/santiagoh1997/weather-api/version-4/testutils"
)

func TestNewJobService(t *testing.T) {
	c := cron.New()
	js := services.NewJobService(nil, nil, c)
	if js.Scheduler == nil {
		t.Errorf("NewWeatherService.Database want %v, got %v", c, nil)
	}
}

func TestNewJob(t *testing.T) {
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)
	ws := &services.WeatherService{
		Database: db,
	}
	c := cron.New()
	l := logger.NewLogger()
	js := services.NewJobService(ws, l, c)
	if apiErr := js.NewJob("Bogotá", "CO"); apiErr != nil {
		t.Errorf("NewJob err = %v, want %v", err, nil)
	}
}
