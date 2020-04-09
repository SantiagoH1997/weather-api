package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-1/models"
	"github.com/santiagoh1997/weather-api/version-1/utils"
)

var apiURL string

func init() {
	apiURL = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=" + beego.AppConfig.String("appid")
}

// WeatherService is the middleman between the controller and the model
type WeatherService struct{}

// Get makes an API call and calls the NewFromResponse
// method from the Weather model
func (ws *WeatherService) Get(city, country string) (*models.Weather, *utils.APIError) {
	url := fmt.Sprintf(apiURL, city, country)
	res, err := http.Get(url)
	if err != nil {
		logErr := fmt.Sprintf("Error while fetching weather: %s", err.Error())
		return nil, utils.NewInternalServerError(logErr)
	}
	defer res.Body.Close()
	var apiRes models.APIResponse
	json.NewDecoder(res.Body).Decode(&apiRes)
	if apiRes.StatusCode != "" {
		statusCode, err := strconv.Atoi(apiRes.StatusCode)
		if err != nil {
			logErr := fmt.Sprintf("Error while converting status code: %s", err.Error())
			return nil, utils.NewInternalServerError(logErr)
		}
		return nil, utils.NewAPIError(statusCode, apiRes.Message)
	}
	var w models.Weather
	w.NewFromResponse(&apiRes)
	return &w, nil
}
