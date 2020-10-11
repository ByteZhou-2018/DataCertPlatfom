package controllers

import (
	"DataCertPlatfom/models"
	"fmt"
	"github.com/astaxie/beego"
)

type SignController struct {
	beego.Controller
}

func (s *SignController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "register.html"
	//s.TplName = "sign_in.html"
	var user models.User
	user.Phone = s.Ctx.Request.Form.Get("phone")
	user.Password =s.Ctx.Request.Form.Get("password")
	fmt.Println(user)
	s.TplName = "index.html"

}
