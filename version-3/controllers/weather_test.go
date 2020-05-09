package controllers_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-3/controllers"
	"github.com/santiagoh1997/weather-api/version-3/logger"
	"github.com/santiagoh1997/weather-api/version-3/models"
	"github.com/santiagoh1997/weather-api/version-3/services"
	"github.com/santiagoh1997/weather-api/version-3/testutils"
)

func TestNewWeatherController(t *testing.T) {
	wc := controllers.NewWeatherController(nil)
	if wc == nil {
		t.Errorf("NewWeatherController = %v, want *WeatherController", wc)
	}
}

func TestGet(t *testing.T) {
	// DB setup
	db, teardown, err := testutils.Setup()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	defer teardown(ctx)

	// Initializing services/controllers
	l := logger.NewLogger()
	ws := services.NewWeatherService(db, l)
	wc := controllers.NewWeatherController(ws)

	// Setting up the handler
	handler := beego.NewControllerRegister()
	handler.Add("/weather", wc, "get:Get")

	services.APIURL = testutils.APIURL

	// Tests
	tests := []struct {
		name             string
		URI              string
		wantStatusCode   int
		wantLocationName string
	}{
		{"Missing city and country", "/weather", 400, ""},
		{"Missing city", "/weather?country=CO", 400, ""},
		{"Missing country", "/weather?city=Bogotá", 400, ""},
		{"No missing params", "/weather?city=Bogotá&country=CO", 200, "Bogotá, CO"},
		{"No missing params", "/weather?city=Paris&country=FR", 200, "Paris, FR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, tt.URI, nil)

			handler.ServeHTTP(w, r)

			statusCode := w.Code
			if statusCode != tt.wantStatusCode {
				t.Fatalf("Get() StatusCode = %v want %v", statusCode, tt.wantStatusCode)
			}

			if tt.wantLocationName != "" {
				res := w.Result()
				defer res.Body.Close()
				resBody, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Fatalf("ReadAll() err = %v want %v", err, nil)
				}

				var w models.Weather
				err = json.Unmarshal(resBody, &w)
				if err != nil {
					t.Fatalf("Unmarshal() err = %v want %v", err, nil)
				}

				if w.LocationName != tt.wantLocationName {
					t.Errorf("Get() LocationName = %v, want %v", w.LocationName, tt.wantLocationName)
				}
			}
		})
	}
}
