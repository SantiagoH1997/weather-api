package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-1/services"
	"github.com/santiagoh1997/weather-api/version-1/utils"
)

// WeatherController connects the request with the model
type WeatherController struct {
	beego.Controller
}

// Get takes a city and a country and
// outputs the corresponding weather
func (wc *WeatherController) Get() {
	var req utils.Request
	req.City = wc.GetString("city")
	req.Country = strings.ToLower(wc.GetString("country"))
	if err := req.ValidateRequest(); err != nil {
		wc.Ctx.Output.SetStatus(err.StatusCode)
		wc.Data["json"] = err
		wc.ServeJSON()
	}

	var ws services.WeatherService
	weather, apiErr := ws.Get(req.City, req.Country)
	if apiErr != nil {
		wc.Ctx.Output.SetStatus(apiErr.StatusCode)
		wc.Data["json"] = apiErr
	} else {
		wc.Data["json"] = weather
	}
	wc.ServeJSON()
}
