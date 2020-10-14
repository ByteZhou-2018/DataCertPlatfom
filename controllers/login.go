package controllers

import (
	"DataCertPlatfom/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "index.html"
}
func (l *LoginController) Post() {

	var user models.User

	err := l.ParseForm(&user)
	if err != nil {
		l.Data["Error"] = err.Error()
		l.TplName = "404.html"
		return
	}
	//fmt.Println(user.Name,user.Password)
	if user.Name == "" || user.Password == "" {
		l.Data["Error"] = "用户名或密码为空！！！"
		l.TplName = "404.html"
		return
	}
	_, err = user.QueryUser()
	if err != nil {
		l.Data["Error"] = err.Error()
		l.TplName = "404.html"
		return
	}
	//l.Data["Username"] = u.Name
	//l.TplName = "home.html"
	l.TplName ="files.html"
}
