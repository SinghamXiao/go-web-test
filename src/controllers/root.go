package controllers

import "github.com/astaxie/beego"

type RootController struct {
	beego.Controller
}

func (this *RootController) Get() {
	uri := this.Ctx.Request.RequestURI
	beego.Info("URI: ", uri)
	this.Ctx.Output.Body([]byte("Go Web Test Success!"))
}
