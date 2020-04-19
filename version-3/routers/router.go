package routers

import (
	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-3/controllers"
)

// MapURLs maps every endpoint with the corresponding controller
func MapURLs(wc *controllers.WeatherController, jc *controllers.JobController) {
	beego.Router("/weather", wc, "get:Get")
	beego.Router("/scheduler/weather", jc, "put:Schedule")
}
