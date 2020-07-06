package main

import (
	"context"
	"fmt"
	"os"

	"github.com/robfig/cron"
	"github.com/santiagoh1997/weather-api/version-5/controllers"
	"github.com/santiagoh1997/weather-api/version-5/datasources/db"
	"github.com/santiagoh1997/weather-api/version-5/logger"
	"github.com/santiagoh1997/weather-api/version-5/pool"
	"github.com/santiagoh1997/weather-api/version-5/routers"
	"github.com/santiagoh1997/weather-api/version-5/services"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/astaxie/beego"
)

const workerCount = 5

func main() {
	l := logger.NewLogger()
	defer l.Sync()

	// DB config
	var mongoDatabase *mongo.Database
	dataSource := os.Getenv("dataSource")
	if dataSource != "json" {
		mongoURI := fmt.Sprintf("mongodb://%s:%s", os.Getenv("mongoHost"), os.Getenv("mongoPort"))
		mongoDB, close, err := db.Open(mongoURI, os.Getenv("mongoDBName"))
		if err != nil {
			l.Error("Error while connecting to the DB")
			panic(err)
		}
		defer close(context.Background())
		l.Info("Connected to the DB")
		mongoDatabase = mongoDB
	}

	// Scheduler
	c := cron.New()
	c.Start()

	// Workers
	collector := pool.StartDispatcher(workerCount) // start up worker pool

	// Services config
	ws := services.NewWeatherService(mongoDatabase, l, &collector)
	js := services.NewJobService(ws, l, c)
	// Controllers config
	wc := controllers.NewWeatherController(ws)
	jc := controllers.NewJobController(js)
	// Router config
	routers.MapURLs(wc, jc)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
