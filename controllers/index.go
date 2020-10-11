package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {
	i.TplName = "index.html"
}