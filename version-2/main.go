package main

import (
	"github.com/santiagoh1997/weather-api/version-2/logger"
	_ "github.com/santiagoh1997/weather-api/version-2/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	defer logger.Log.Sync()
	beego.Run()
}
