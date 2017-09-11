package routers

import (
	"github.com/astaxie/beego"
	"controllers"
)

func init() {
	// Register routers.
	beego.Router("/", &controllers.RootController{})
}
