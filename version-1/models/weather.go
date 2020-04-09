package models

import (
	"fmt"
	"time"
)

// Weather is the struct for the response
type Weather struct {
	LocationName   string    `json:"location_name"`
	Temperature    string    `json:"temperature"`
	Wind           string    `json:"wind"`
	Cloudiness     string    `json:"cloudiness"`
	Pressure       string    `json:"pressure"`
	Humidity       string    `json:"humidity"`
	Sunrise        string    `json:"sunrise"`
	Sunset         string    `json:"sunset"`
	GeoCoordinates []float32 `json:"geo_coordinates"`
	RequestedTime  string    `json:"requested_time"`
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
