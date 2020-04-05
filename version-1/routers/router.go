package routers

import (
	"github.com/santiagoh1997/weather-api/version-1/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/weather", &controllers.WeatherController{}, "get:Get")
}
