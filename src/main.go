package main

import (
	_ "logs"
	_ "routers"
	"github.com/astaxie/beego"
)

var version string = beego.AppConfig.String("version")
var host string = beego.AppConfig.String("host")
var port string = beego.AppConfig.String("port")

func main() {
	appName := beego.BConfig.AppName
	beego.Notice(appName, version)
	beego.Run(host + ":" + port)
}
