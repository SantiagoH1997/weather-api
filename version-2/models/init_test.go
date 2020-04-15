package models_test

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"path/filepath"

// 	"go.mongodb.org/mongo-driver/mongo"

// 	"github.com/astaxie/beego"
// 	"github.com/santiagoh1997/weather-api/version-2/db"
// )

// var mongoURI string

// func init() {
// 	pwd, err := os.Getwd()
// 	if err != nil {
// 		panic(err)
// 	}
// 	beego.TestBeegoInit(filepath.Dir(pwd))
// 	mongoURI = beego.AppConfig.String("mongoTestHost")
// }

// func setup() (*mongo.Client, func(ctx context.Context) error, error) {
// 	database, close, err := db.Open(mongoURI)
// 	if err != nil {
// 		log.Fatal("Couldn't connect to MongoDB")
// 	}
// 	return database.MongoClient, close, nil
// }
