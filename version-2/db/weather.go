package db

import (
	"context"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/santiagoh1997/weather-api/version-2/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	mongoURI string
	dbName   string
	// Database is the MongoDB database
	Database          *mongo.Database
	weatherCollection = "weather"
	indexKey          = "location_name"
)

func init() {
	if beego.BConfig.RunMode == "dev" {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", beego.AppConfig.String("mongoHost"), beego.AppConfig.String("mongoPort"))
	}
	dbName = beego.AppConfig.String("mongoDBName")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Log.Error(err)
		panic(err)
	}
	Database = client.Database(dbName)
	index := mongo.IndexModel{}
	index.Keys = bsonx.Doc{{Key: indexKey, Value: bsonx.String("text")}}
	index.Options = options.Index().SetUnique(true)
	_, err = Database.Collection(weatherCollection).Indexes().CreateOne(ctx, index)
	if err != nil {
		logger.Log.Error(err)
		panic(err)
	}
	logger.Log.Info("Connected to MongoDB")
}
