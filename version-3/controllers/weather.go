package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-3/services"
	"github.com/santiagoh1997/weather-api/version-3/utils"
)

// WeatherController connects the request with the model
type WeatherController struct {
	beego.Controller
	Service *services.WeatherService
}

// NewWeatherController returns a pointer to a WeatherController
func NewWeatherController(ws *services.WeatherService) *WeatherController {
	return &WeatherController{
		Service: ws,
	}
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
		return
	}
	weather, apiErr := wc.Service.Get(req.City, req.Country)
	if apiErr != nil {
		wc.Ctx.Output.SetStatus(apiErr.StatusCode)
		wc.Data["json"] = apiErr
	} else {
		wc.Data["json"] = weather
	}
	wc.ServeJSON()
}
