package testutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-4/datasources/db"
	"github.com/santiagoh1997/weather-api/version-4/testdata"
)

var (
	mongoURI string
	dbName   string
	// APIURL is the URL for the third party API
	APIURL     string
	collection = "weather"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	beego.TestBeegoInit(filepath.Dir(pwd))
	mongoURI = fmt.Sprintf("mongodb://%s:%s", beego.AppConfig.String("mongoTestHost"), beego.AppConfig.String("mongoPort"))
	dbName = beego.AppConfig.String("mongoTestDBName")
	APIURL = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=" + beego.AppConfig.String("appid")
}

// Setup connects to the DB and seeds it.
// It returns a connection to the DB, a function that closes the connection, and an error
func Setup() (*mongo.Database, func(ctx context.Context) error, error) {
	database, close, err := db.Open(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Couldn't connect to MongoDB: %v", err.Error())
	}
	ctx := context.Background()
	if _, err := database.Collection(collection).DeleteMany(ctx, bson.M{}); err != nil {
		log.Fatalf("Error while deleting records from the DB: %v", err.Error())
	}
	if err := seed(ctx, database); err != nil {
		log.Fatalf("Error while seeding the DB: %v", err.Error())
	}
	return database, close, nil
}

func seed(ctx context.Context, db *mongo.Database) error {
	weathers := testdata.TestWeathers
	weathersInterface := make([]interface{}, len(weathers))
	for i, v := range weathers {
		weathersInterface[i] = v
	}
	_, err := db.Collection(collection).InsertMany(ctx, weathersInterface)
	return err
}
