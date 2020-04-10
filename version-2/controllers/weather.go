package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-2/services"
)

// WeatherController connects the request with the model
type WeatherController struct {
	beego.Controller
}

// Get takes a city and a country and
// outputs the corresponding weather
func (wc *WeatherController) Get() {
	city := wc.GetString("city")
	country := strings.ToLower(wc.GetString("country"))
	if city != "" && country != "" {
		var ws services.WeatherService
		weather, apiErr := ws.Get(city, country)
		if apiErr != nil {
			wc.Ctx.Output.SetStatus(apiErr.StatusCode)
			wc.Data["json"] = apiErr
		} else {
			wc.Data["json"] = weather
		}
		wc.ServeJSON()
	}
}
