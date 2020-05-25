package routers

import (
	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-4/controllers"
)

// MapURLs maps every endpoint with the corresponding controller
func MapURLs(wc *controllers.WeatherController, jc *controllers.JobController) {
	// swagger:route GET /weather weather getWeather
	// Returns a weather report given a city and a country
	// responses:
	// 	200: weather
	// 	404: notFoundError
	// 	401: badRequestError
	beego.Router("/weather", wc, "get:Get")
	// swagger:route PUT /scheduler/weather job scheduleJob
	// Schedules a job to be performed every 1 hour
	// responses:
	// 	202: jobScheduledResponse
	// 	401: badRequestError
	beego.Router("/scheduler/weather", jc, "put:Schedule")
}
