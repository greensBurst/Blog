package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Data["isHome"] = true
	this.TplName = "home.html"
	this.Data["isLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["topics"] = topics
	}
}
