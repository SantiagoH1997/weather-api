package main

import (
	"context"
	"fmt"

	"github.com/robfig/cron"
	"github.com/santiagoh1997/weather-api/version-4/controllers"
	"github.com/santiagoh1997/weather-api/version-4/datasources/db"
	"github.com/santiagoh1997/weather-api/version-4/logger"
	"github.com/santiagoh1997/weather-api/version-4/routers"
	"github.com/santiagoh1997/weather-api/version-4/services"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/astaxie/beego"
)

func main() {
	l := logger.NewLogger()
	defer l.Sync()

	// DB config
	var mongoDatabase *mongo.Database
	dataSource := beego.AppConfig.String("dataSource")
	if dataSource != "json" {
		var mongoURI string
		if beego.BConfig.RunMode != "prod" {
			mongoURI = fmt.Sprintf("mongodb://%s:%s", beego.AppConfig.String("mongoHost"), beego.AppConfig.String("mongoPort"))
		}
		mongoDB, close, err := db.Open(mongoURI, beego.AppConfig.String("mongoDBName"))
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

	// Services config
	ws := services.NewWeatherService(mongoDatabase, l)
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
