package main

import (
	"context"
	"fmt"

	"github.com/robfig/cron"
	"github.com/santiagoh1997/weather-api/version-3/controllers"
	"github.com/santiagoh1997/weather-api/version-3/db"
	"github.com/santiagoh1997/weather-api/version-3/logger"
	"github.com/santiagoh1997/weather-api/version-3/routers"
	"github.com/santiagoh1997/weather-api/version-3/services"

	"github.com/astaxie/beego"
)

func main() {
	l := logger.NewLogger()
	defer l.Sync()
	// DB config
	mongoURI := fmt.Sprintf("mongodb://%s:%s", beego.AppConfig.String("mongoHost"), beego.AppConfig.String("mongoPort"))
	mongoDB, close, err := db.Open(mongoURI, beego.AppConfig.String("mongoDBName"))
	if err != nil {
		l.Error("Error while connecting to the DB")
		panic(err)
	}
	defer close(context.Background())
	l.Info("Connected to the DB")
	// Scheduler

	c := cron.New()
	c.Start()
	// Services config
	ws := services.NewWeatherService(mongoDB, l)
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
