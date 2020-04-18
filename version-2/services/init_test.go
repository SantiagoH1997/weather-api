package services_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-2/db"
	"github.com/santiagoh1997/weather-api/version-2/testdata"
)

var (
	mongoURI   string
	dbName     string
	apiURL     string
	collection = "weather"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	beego.TestBeegoInit(filepath.Dir(pwd))
	mongoURI = beego.AppConfig.String("mongoTestHost")
	mongoURI = fmt.Sprintf("mongodb://%s:%s", beego.AppConfig.String("mongoTestHost"), beego.AppConfig.String("mongoPort"))
	dbName = beego.AppConfig.String("mongoTestDBName")
	apiURL = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&units=metric&appid=" + beego.AppConfig.String("appid")
}

func setup() (*mongo.Database, func(ctx context.Context) error, error) {
	database, close, err := db.Open(mongoURI, dbName)
	if err != nil {
		log.Fatalf("Couldn't connect to MongoDB: %v", err.Error())
	}
	ctx := context.Background()
	if _, err := database.Collection(collection).DeleteMany(ctx, bson.M{}); err != nil {
		log.Fatalf("Error while deleting records from the DB: %v", err.Error())
	}
	if err := seed(database, ctx); err != nil {
		log.Fatalf("Error while seeding the DB: %v", err.Error())
	}
	return database, close, nil
}

func seed(db *mongo.Database, ctx context.Context) error {
	weathers := testdata.TestWeathers
	weathersInterface := make([]interface{}, len(weathers))
	for i, v := range weathers {
		weathersInterface[i] = v
	}
	_, err := db.Collection(collection).InsertMany(ctx, weathersInterface)
	return err
}
