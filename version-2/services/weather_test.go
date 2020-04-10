package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	testCity         = "Bogota"
	testCountry      = "co"
	testLocationName = "Bogotá, CO"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	beego.TestBeegoInit(filepath.Dir(pwd))
	apiURL = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=" + beego.AppConfig.String("appid")
}

func TestGet(t *testing.T) {
	Convey("Given a city and a country", t, func() {
		var ws WeatherService
		Convey("It should make the corresponding API call", func() {
			w, err := ws.Get(testCity, testCountry)
			So(err, ShouldBeNil)
			So(w.LocationName, ShouldEqual, testLocationName)
		})
		Convey("It should return a 404 status code if the city/country is not found", func() {
			w, err := ws.Get("abc", "123")
			So(w, ShouldBeNil)
			So(err.StatusCode, ShouldEqual, 404)
		})
	})
}
