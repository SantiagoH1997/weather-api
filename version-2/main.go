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
	// DB config
	var mongoURI string
	if beego.BConfig.RunMode != "prod" {
		mongoURI = fmt.Sprintf("mongodb://%s:%s", beego.AppConfig.String("mongoHost"), beego.AppConfig.String("mongoPort"))
	}
	db, close, err := db.Open(mongoURI, beego.AppConfig.String("mongoDBName"))
	if err != nil {
		logger.Log.Error("Error while connecting to the DB")
		panic(err)
	}
	defer close(context.Background())
	logger.Log.Info("Connected to the DB")

	// Services config
	ws := services.NewWeatherService(db)
	// Controllers config
	wc := controllers.NewWeatherController(ws)

	// Router config
	routers.MapURLs(wc)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	defer logger.Log.Sync()
	beego.Run()
}
