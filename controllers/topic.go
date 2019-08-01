package controllers

import (
	"blog/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["isLogin"] = checkAccount(this.Ctx)
	this.Data["isTopic"] = true
	this.TplName = "topic.html"

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["topics"] = topics
	}
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 301)
		return
	}

	title := this.Input().Get("title")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 301)
}

func (this *TopicController) Add() {
	this.Data["isLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	this.Data["isLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_view.html"

	mp := this.Ctx.Input.Params() // /view/【“0”】/【“1”】
	topic, err := models.GetTopic(mp["0"])
	if err != nil {
		beego.Error(err)
		this.Redirect("/topic", 301)
		return
	}
	this.Data["topic"] = topic
	this.Data["tid"] = mp["0"]
}

func (this *TopicController) Modify() {
	this.Data["isLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic_modify.html"
	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/topic", 301)
		return
	}
	this.Data["topic"] = topic
	this.Data["tid"] = tid
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 301)
		return
	}

	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/", 301)
}
