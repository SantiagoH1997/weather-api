package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-2/models"
	"github.com/santiagoh1997/weather-api/version-2/utils"
)

var apiURL string

const expirationTime = 300 * time.Second

func init() {
	apiURL = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=" + beego.AppConfig.String("appid")
}

// WeatherService is the middleman between the controller and the model
type WeatherService struct{}

// Get retrieves an existing Weather from the DB.
// If it's not in the DB, it fetches a new weather report from the API and saves it.
// If the Weather is old, it updates it.
func (ws *WeatherService) Get(city, country string) (*models.Weather, *utils.APIError) {
	var w models.Weather
	err := w.GetByLocation(fmt.Sprintf("%s, %s", city, country))
	hasExpired := time.Since(w.ModifiedAt.Time()) > expirationTime
	if err != nil || hasExpired {
		url := fmt.Sprintf(apiURL, city, country)
		if err != nil {
			if err.StatusCode != http.StatusNotFound {
				return nil, err
			}
			if err.StatusCode == http.StatusNotFound {
				apiRes, apiErr := ws.fetchWeather(url)
				if apiErr != nil {
					return nil, apiErr
				}
				w.NewFromResponse(apiRes)
				fmt.Println("Saving...")
				w.Save()
			}
		} else {
			apiRes, apiErr := ws.fetchWeather(url)
			if apiErr != nil {
				return nil, apiErr
			}
			w.NewFromResponse(apiRes)
			fmt.Println("Updating...")
			w.Update()
		}
	}
	return &w, nil
}

// fetchWeather fetches a weather report from the API
func (ws *WeatherService) fetchWeather(url string) (*models.APIResponse, *utils.APIError) {
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
	return &apiRes, nil
}
