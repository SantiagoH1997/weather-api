package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Weather is the response for GET /weather
// swagger:response weather
// in: body
type Weather struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	LocationName   string             `json:"location_name" bson:"location_name"`
	Temperature    string             `json:"temperature" bson:"temperature"`
	Wind           string             `json:"wind" bson:"wind"`
	Cloudiness     string             `json:"cloudiness" bson:"cloudiness"`
	Pressure       string             `json:"pressure" bson:"pressure"`
	Humidity       string             `json:"humidity" bson:"humidity"`
	Sunrise        string             `json:"sunrise" bson:"sunrise"`
	Sunset         string             `json:"sunset" bson:"sunset"`
	GeoCoordinates []float32          `json:"geo_coordinates" bson:"geo_coordinates"`
	RequestedTime  string             `json:"requested_time" bson:"requested_time"`
	ModifiedAt     primitive.DateTime `json:"-" bson:"modified_at"`
	Score          float32            `json:"-" bson:"score"`
}

// NewWeatherFromResponse populates the fields of
// a Weather struct with values from the API response
func NewWeatherFromResponse(ar *APIResponse) *Weather {
	return &Weather{
		LocationName:   fmt.Sprintf("%s, %s", ar.City, ar.Sys.Country),
		Temperature:    fmt.Sprintf("%v °C", ar.Main.Temperature),
		Wind:           fmt.Sprintf("%v m/s", ar.Wind.Speed),
		Cloudiness:     fmt.Sprintf("%d%%", ar.Clouds.Cloudiness),
		Pressure:       fmt.Sprintf("%d hpa", ar.Main.Pressure),
		Humidity:       fmt.Sprintf("%d%%", ar.Main.Humidity),
		Sunrise:        time.Unix(ar.Sys.Sunrise, 0).Format("03:04"),
		Sunset:         time.Unix(ar.Sys.Sunset, 0).Format("03:04"),
		GeoCoordinates: []float32{ar.Coord.Lat, ar.Coord.Lon},
		RequestedTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
}
