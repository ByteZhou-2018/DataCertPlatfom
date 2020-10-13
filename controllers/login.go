package controllers

import (
	"DataCertPlatfom/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "index.html"
}
func (l *LoginController)Post()  {
	
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		fmt.Println("登录页面解析用户数据错误！",err.Error())
	}
	u,err := user.QueryUser()
	if err != nil {
		l.Data["Error"] = err.Error()
		l.TplName = "404.html"
		return
	}
	l.Data["Username"] = u.Name
	//l.Ctx.WriteString("登录成功 ! 已到达用户主页面!")
	l.TplName ="home.html"
}
