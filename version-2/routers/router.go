package routers

import (
	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-2/controllers"
)

// MapURLs maps every endpoint with the corresponding controller
func MapURLs(wc *controllers.WeatherController) {
	beego.Router("/weather", wc, "get:Get")
}
