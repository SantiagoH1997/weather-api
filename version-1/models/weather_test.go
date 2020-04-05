package models

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/astaxie/beego"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	testAPIResponse = `{"coord":{"lon":-74.08,"lat":4.61},"weather":[
{"id":803,"main":"Clouds","description":"broken clouds","icon":"04d"}],
"base":"stations","main":{"temp":292.15,"feels_like":289.39,"temp_min":292.15,
"temp_max":292.15,"pressure":1026,"humidity":52},"visibility":10000,"wind":{"speed":3.6,"deg":30},
"clouds":{"all":75},"dt":1586116410,"sys":{"type":1,"id":8582,"country":"CO","sunrise":1586083995,
"sunset":1586127843},"timezone":-18000,"id":3688689,"name":"Bogotá","cod":"200"}`
	testLocationName = "Bogotá, CO"
	testTemperature  = "292.15 °C"
	testWind         = "3.6 m/s"
	testCloudiness   = "75%"
	testPressure     = "1026 hpa"
	testHumidity     = "52%"
	testSunrise      = "07:53"
	testSunset       = "08:04"
	testCity         = "Bogota"
	testCountry      = "co"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	beego.TestBeegoInit(filepath.Dir(pwd))
	apiURL = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=" + beego.AppConfig.String("appid")
}

func TestNewFromResponse(t *testing.T) {
	var apiRes apiResponse
	err := json.Unmarshal([]byte(testAPIResponse), &apiRes)
	if err != nil {
		panic(err)
	}

	Convey("Given an API response", t, func() {
		Convey("It should return a new Weather struct with the correct fields", func() {
			var w Weather
			w.NewFromResponse(&apiRes)
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

func TestGet(t *testing.T) {
	Convey("Given a city and a country", t, func() {
		Convey("It should return a new Weather struct with the correct location name", func() {
			w, err := Get(testCity, testCountry)
			So(err, ShouldBeNil)
			So(w.LocationName, ShouldEqual, testLocationName)
		})
		// Convey("It should return a 404 status code if the city/country is not found", func() {
		// 	w, err := Get("abc", "123")
		// 	So(w, ShouldBeNil)
		// 	So(err.StatusCode, ShouldEqual, 404)
		// })
	})
}
