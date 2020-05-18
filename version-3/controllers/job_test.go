package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	"github.com/robfig/cron"
	"github.com/santiagoh1997/weather-api/version-3/controllers"
	"github.com/santiagoh1997/weather-api/version-3/services"
)

func TestNewJobController(t *testing.T) {
	jc := controllers.NewJobController(nil)
	if jc == nil {
		t.Errorf("NewJobController = %v, want *JobController", jc)
	}
}

func TestSchedule(t *testing.T) {

	c := cron.New()

	// Initializing services/controllers
	ws := services.NewWeatherService(nil, nil)
	js := services.NewJobService(ws, nil, c)
	jc := controllers.NewJobController(js)

	// Setting up the handler
	handler := beego.NewControllerRegister()
	handler.Add("/scheduler/weather", jc, "get:Schedule")

	// Tests
	tests := []struct {
		name string
		URI  string
		want int
	}{
		{"Missing city and country", "/scheduler/weather", 400},
		{"Missing city", "/scheduler/weather?country=CO", 400},
		{"Missing country", "/scheduler/weather?city=Bogotá", 400},
		{"No missing params", "/scheduler/weather?city=Bogotá&country=CO", 202},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, tt.URI, nil)

			handler.ServeHTTP(w, r)

			got := w.Code
			if got != tt.want {
				t.Fatalf("Schedule() StatusCode = %v want %v", got, tt.want)
			}
		})
	}
}
