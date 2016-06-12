package main

import (
	_ "skitta.com/goweb/sinanewsAPI/docs"
	_ "skitta.com/goweb/sinanewsAPI/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
