package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Weather is the struct for the response
type Weather struct {
	ID             primitive.ObjectID `json:"-" bson:"_id"`
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
