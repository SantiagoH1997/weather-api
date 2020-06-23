package models_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/santiagoh1997/weather-api/version-5/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	id          primitive.ObjectID
	dateTime    primitive.DateTime
	testWeather = &models.Weather{
		LocationName:   "Bogotá, CO",
		Temperature:    "292.15 °C",
		Wind:           "3.6 m/s",
		Cloudiness:     "75%",
		Pressure:       "1026 hpa",
		Humidity:       "52%",
		Sunrise:        "07:53",
		Sunset:         "08:04",
		GeoCoordinates: []float32{4.61, -74.08},
		RequestedTime:  "",
	}
)

func TestNewFromResponse(t *testing.T) {
	var apiRes models.APIResponse
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(filepath.Join(pwd, "..", "testdata", "response.json"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(data), &apiRes); err != nil {
		panic(err)
	}
	got := models.NewWeatherFromResponse(&apiRes)

	tests := []struct {
		name string
		got  interface{}
		want interface{}
	}{
		{"LocationName", got.LocationName, testWeather.LocationName},
		{"Temperature", got.Temperature, testWeather.Temperature},
		{"Wind", got.Wind, testWeather.Wind},
		{"Cloudiness", got.Cloudiness, testWeather.Cloudiness},
		{"Pressure", got.Pressure, testWeather.Pressure},
		{"Humidity", got.Humidity, testWeather.Humidity},
		{"Sunrise", got.Sunrise, testWeather.Sunrise},
		{"Sunset", got.Sunset, testWeather.Sunset},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("NewFromResponse %s got %v want %v", tt.name, tt.got, tt.want)
			}
		})
	}

}
