package controllers

import (
	"DataCertPlatfom/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.Data["Username"] = "游客32098178765"
	l.TplName = "login.html"
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
	u, err := user.QueryUser()
	if err != nil {
		l.Data["Error"] = err.Error()
		l.TplName = "404.html"
		return
	}
	l.Data["Username"] = u.Name
	l.Data["Phone"] = u.Phone
	//l.TplName = "home.html"
	l.TplName ="files.html"
}
