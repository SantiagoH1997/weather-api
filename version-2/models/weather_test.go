package models_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/santiagoh1997/weather-api/version-2/models"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	testLocationName = "Bogotá, CO"
	testTemperature  = "292.15 °C"
	testWind         = "3.6 m/s"
	testCloudiness   = "75%"
	testPressure     = "1026 hpa"
	testHumidity     = "52%"
	testSunrise      = "07:53"
	testSunset       = "08:04"
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

	Convey("Given an API response", t, func() {
		Convey("It should return a new Weather struct with the correct fields", func() {
			w := models.NewWeatherFromResponse(&apiRes)
			So(w.LocationName, ShouldEqual, testLocationName)
			So(w.Temperature, ShouldEqual, testTemperature)
			So(w.Wind, ShouldEqual, testWind)
			So(w.Cloudiness, ShouldEqual, testCloudiness)
			So(w.Pressure, ShouldEqual, testPressure)
			So(w.Humidity, ShouldEqual, testHumidity)
			So(w.Sunrise, ShouldEqual, testSunrise)
			So(w.Sunset, ShouldEqual, testSunset)
		})
	})
}
