package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/santiagoh1997/weather-api/version-5/models"
	"github.com/santiagoh1997/weather-api/version-5/pool"
	"github.com/santiagoh1997/weather-api/version-5/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// APIURL is exported for testing purposes...
var (
	APIURL string
)

const (
	expirationTime   = 300 * time.Second
	collection       = "weather"
	timeout          = time.Second * 10
	minimumTextScore = 1.5
)

func init() {
	APIURL = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=" + os.Getenv("appid")
}

// WeatherService interacts with the persistance layer
// and fetchs new weather reports from the API
type WeatherService struct {
	Database  *mongo.Database
	logger    *zap.SugaredLogger
	collector *pool.Collector
}

// NewWeatherService returns a pointer to a WeatherService
func NewWeatherService(db *mongo.Database, l *zap.SugaredLogger, c *pool.Collector) *WeatherService {
	return &WeatherService{
		db,
		l,
		c,
	}
}

// Get retrieves an existing Weather from the DB.
// If it's not in the DB, it fetches a new weather report from the API and saves it.
// If the Weather is old, it updates it.
func (ws *WeatherService) Get(city, country string) (*models.Weather, *utils.APIError) {
	location := fmt.Sprintf("%s, %s", city, country)

	if ws.Database == nil {
		return ws.GetFromJSON(location)
	}

	weather, err := ws.GetByLocation(location)
	url := fmt.Sprintf(APIURL, city, country)
	if err != nil {
		if err.StatusCode == http.StatusNotFound {
			apiRes, apiErr := ws.FetchWeather(url)
			if apiErr != nil {
				return nil, apiErr
			}
			weather = models.NewWeatherFromResponse(apiRes)
			ws.Save(weather)
			return weather, nil
		}
	}

	// If weather report has expired
	if time.Since(weather.ModifiedAt.Time()) > expirationTime {
		apiRes, apiErr := ws.FetchWeather(url)
		if apiErr != nil {
			return nil, apiErr
		}
		weather = models.NewWeatherFromResponse(apiRes)
		ws.Update(weather)
	}

	return weather, nil
}

// GetByLocation retrieves a weather report from the DB
func (ws *WeatherService) GetByLocation(location string) (*models.Weather, *utils.APIError) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	search := bson.M{
		"$text": bson.M{
			"$search": location,
		},
	}
	options := options.FindOne().SetProjection(bson.M{
		"score": bson.M{
			"$meta": "textScore",
		},
	})

	var w models.Weather
	err := ws.Database.Collection(collection).FindOne(ctx, search, options).Decode(&w)
	if err != nil || w.Score < minimumTextScore {
		if err == mongo.ErrNoDocuments || w.Score < minimumTextScore {
			return nil, utils.NewNotFound("No results found")
		}
		ws.logger.Error(fmt.Sprintf("Error while retrieving weather: %s", err.Error()))
		return nil, utils.NewInternalServerError()
	}

	return &w, nil
}

// Save saves a weather response to the database
func (ws *WeatherService) Save(w *models.Weather) (*mongo.InsertOneResult, *utils.APIError) {
	w.ModifiedAt = primitive.NewDateTimeFromTime(time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := ws.Database.Collection(collection).InsertOne(ctx, w)
	if err != nil {
		ws.logger.Error(fmt.Sprintf("Error while saving weather: %s", err.Error()))
		return nil, utils.NewInternalServerError()
	}
	return res, nil
}

// Update updates an existing weather in the database
func (ws *WeatherService) Update(w *models.Weather) (*mongo.UpdateResult, *utils.APIError) {
	w.ModifiedAt = primitive.NewDateTimeFromTime(time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	filter := bson.M{
		"location_name": w.LocationName,
	}
	update := bson.M{
		"$set": w,
	}
	res, err := ws.Database.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		ws.logger.Error(fmt.Sprintf("Error while updating weather: %s", err.Error()))
		return nil, utils.NewInternalServerError()
	}
	return res, nil
}

// FetchWeather fetches a weather report from the API
func (ws *WeatherService) FetchWeather(url string) (*models.APIResponse, *utils.APIError) {
	rc := make(chan *http.Response)
	ec := make(chan error)
	ws.collector.Work <- pool.Work{URL: url, ResponseChannel: rc, ErrorChannel: ec}

	var res *http.Response
	select {
	case r := <-rc:
		res = r
	case err := <-ec:
		ws.logger.Error(fmt.Sprintf("Error while fetching weather: %s", err.Error()))
		return nil, utils.NewInternalServerError()
	}
	defer res.Body.Close()

	var apiRes models.APIResponse
	json.NewDecoder(res.Body).Decode(&apiRes)
	if apiRes.StatusCode != "" {
		statusCode, err := strconv.Atoi(apiRes.StatusCode)
		if err != nil {
			ws.logger.Error(fmt.Sprintf("Error while converting status code: %s", err.Error()))
			return nil, utils.NewInternalServerError()
		}
		return nil, utils.NewAPIError(statusCode, apiRes.Message)
	}

	return &apiRes, nil
}

// GetFromJSON gets a weather report from the weather.json file
func (ws *WeatherService) GetFromJSON(location string) (*models.Weather, *utils.APIError) {
	path := filepath.Join("datasources", "jsondata", "weather.json")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		ws.logger.Error(fmt.Sprintf("Error while opening json file: %s", err.Error()))
		return nil, utils.NewInternalServerError()
	}

	var weatherData []models.Weather

	err = json.Unmarshal(file, &weatherData)
	if err != nil {
		ws.logger.Error(fmt.Sprintf("Error while unmarshalling json file: %s", err.Error()))
		return nil, utils.NewInternalServerError()
	}

	for _, weather := range weatherData {
		if weather.LocationName == location {
			return &weather, nil
		}
	}

	return nil, utils.NewNotFound("No results found")
}
