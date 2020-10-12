package controllers

import (
	"DataCertPlatfom/db_mysql"
	"DataCertPlatfom/models"
	"fmt"
	"github.com/astaxie/beego"
)

type SignController struct {
	beego.Controller
}

func (s *SignController)Get()  {
	s.TplName = "sign_in.html"

}
func (s *SignController) Post() {
	//fmt.Println("hello sign")
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "register.html"
	//s.TplName = "sign_in.html"
	var user models.User
	err := s.ParseForm(&user)
	if err != nil{
		 s.Ctx.WriteString("数据解析失败，请重试！")
	}
	fmt.Println(user)

	_,err = db_mysql.Inseret(user)
	if err  != nil{
		fmt.Println(err.Error())
		s.Ctx.WriteString("用户注册失败，请返回页面重试！")

	}


	s.TplName = "index.html"

}
