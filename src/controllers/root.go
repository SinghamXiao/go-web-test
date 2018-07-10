package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	uri := this.Ctx.Request.RequestURI
	beego.Info("URI: ", uri)
	this.Ctx.Output.Body([]byte("Go Web Test Success!"))
}
