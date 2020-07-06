package main

import (
	"context"
	"fmt"

	"github.com/santiagoh1997/weather-api/version-2/controllers"
	"github.com/santiagoh1997/weather-api/version-2/db"
	"github.com/santiagoh1997/weather-api/version-2/logger"
	"github.com/santiagoh1997/weather-api/version-2/routers"
	"github.com/santiagoh1997/weather-api/version-2/services"

	"github.com/astaxie/beego"
)

func main() {
	l := logger.NewLogger()
	// DB config
	mongoURI := fmt.Sprintf("mongodb://%s:%s", beego.AppConfig.String("mongoHost"), beego.AppConfig.String("mongoPort"))
	mongoDB, close, err := db.Open(mongoURI, beego.AppConfig.String("mongoDBName"))
	if err != nil {
		l.Error("Error while connecting to the DB")
		panic(err)
	}
	defer close(context.Background())
	l.Info("Connected to the DB")

	// Services config
	ws := services.NewWeatherService(mongoDB, l)
	// Controllers config
	wc := controllers.NewWeatherController(ws)

	// Router config
	routers.MapURLs(wc)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	defer l.Sync()
	beego.Run()
}
