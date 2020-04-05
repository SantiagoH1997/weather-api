package controllers

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-1/logger"
	"github.com/santiagoh1997/weather-api/version-1/models"
)

const internalServerErrorMessage = "Internal server error"

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
		weather, apiErr := models.Get(city, country)
		if apiErr != nil {
			if apiErr.StatusCode >= http.StatusInternalServerError {
				wc.Ctx.Output.SetStatus(apiErr.StatusCode)
				logger.Log.Errorf("Error while fetching weather: %s", apiErr.Message)
				apiErr.Message = internalServerErrorMessage
				wc.Data["json"] = apiErr
			} else {
				wc.Ctx.Output.SetStatus(apiErr.StatusCode)
				wc.Data["json"] = apiErr
			}
		} else {
			wc.Data["json"] = weather
		}
		wc.ServeJSON()
	}
}
