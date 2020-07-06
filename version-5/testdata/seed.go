package testdata

import (
	"time"

	"github.com/santiagoh1997/weather-api/version-5/models"
)

var (
	sun           string
	requestedTime string
	// TestWeathers is a slice of Weathers to seed the DB
	TestWeathers []models.Weather
	// SampleWeather is a sample Weather for testing purposes
	SampleWeather = models.Weather{
		LocationName:  "Tokyo, JP",
		Temperature:   "10 °C",
		Wind:          "100 m/s",
		Cloudiness:    "10%%",
		Pressure:      "100 hpa",
		Humidity:      "10%%",
		Sunrise:       sun,
		Sunset:        sun,
		RequestedTime: requestedTime,
	}
)

func init() {
	sun = time.Now().Format("03:04")
	requestedTime = time.Now().Format("2006-01-02 15:04:05")
	TestWeathers = []models.Weather{
		models.Weather{
			LocationName:  "Bogotá, CO",
			Temperature:   "10 °C",
			Wind:          "100 m/s",
			Cloudiness:    "10%%",
			Pressure:      "100 hpa",
			Humidity:      "10%%",
			Sunrise:       sun,
			Sunset:        sun,
			RequestedTime: requestedTime,
		},
		models.Weather{
			LocationName:  "Paris, FR",
			Temperature:   "10 °C",
			Wind:          "100 m/s",
			Cloudiness:    "10%%",
			Pressure:      "100 hpa",
			Humidity:      "10%%",
			Sunrise:       sun,
			Sunset:        sun,
			RequestedTime: requestedTime,
		},
	}
}
