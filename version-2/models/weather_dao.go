package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/santiagoh1997/weather-api/version-2/db"
	"github.com/santiagoh1997/weather-api/version-2/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collection       = "weather"
	timeout          = time.Second * 10
	minimumTextScore = 1.5
)

// GetByLocation retrieves a weathr report from the DB
func (w *Weather) GetByLocation(location string) *utils.APIError {
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
	err := db.Database.Collection(collection).FindOne(ctx, search, options).Decode(w)
	if err != nil || w.Score < minimumTextScore {
		if err == mongo.ErrNoDocuments || w.Score < minimumTextScore {
			return utils.NewNotFound("No results found")
		}
		return utils.NewInternalServerError(fmt.Sprintf("Error while retrieving weather: %s", err.Error()))
	}
	return nil
}

// Save saves a weather response to the database
func (w *Weather) Save() *utils.APIError {
	w.ModifiedAt = primitive.NewDateTimeFromTime(time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	_, err := db.Database.Collection(collection).InsertOne(ctx, w)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Error while saving weather: %s", err.Error()))
	}
	return nil
}

// Update updates an existing weather in the database
func (w *Weather) Update() *utils.APIError {
	fmt.Println(w)
	w.ModifiedAt = primitive.NewDateTimeFromTime(time.Now())
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	filter := bson.M{
		"location_name": w.LocationName,
	}
	update := bson.M{
		"$set": w,
	}
	_, err := db.Database.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return utils.NewInternalServerError(fmt.Sprintf("Error while updating weather: %s", err.Error()))
	}
	return nil
}

// NewFromResponse populates the fields of
// a Weather struct with values from the API response
func (w *Weather) NewFromResponse(ar *APIResponse) {
	w.LocationName = fmt.Sprintf("%s, %s", ar.City, ar.Sys.Country)
	w.Temperature = fmt.Sprintf("%v °C", ar.Main.Temperature)
	w.Wind = fmt.Sprintf("%v m/s", ar.Wind.Speed)
	w.Cloudiness = fmt.Sprintf("%d%%", ar.Clouds.Cloudiness)
	w.Pressure = fmt.Sprintf("%d hpa", ar.Main.Pressure)
	w.Humidity = fmt.Sprintf("%d%%", ar.Main.Humidity)
	w.Sunrise = time.Unix(ar.Sys.Sunrise, 0).Format("03:04")
	w.Sunset = time.Unix(ar.Sys.Sunset, 0).Format("03:04")
	w.GeoCoordinates = []float32{ar.Coord.Lat, ar.Coord.Lon}
	w.RequestedTime = time.Now().Format("2006-01-02 15:04:05")
}
