package controllers

import (
	"DataCertPlatfom/models"
	"github.com/astaxie/beego"
)

type SignController struct {
	beego.Controller
}

func (s *SignController)Get()  {
	s.TplName = "sign_in.html"

}
func (s *SignController) Post() {
	// 1、解析前端传递过来的数据
	var user models.User

	err := s.ParseForm(&user)
	if err != nil{
		s.Ctx.WriteString("注册页面 用户数据解析数据错误")
		return
	}
	if user.Name == "" || user.Password == "" {
			s.TplName = "sign_in.html"
			return
	}
	//2、将解析到的数保存到数据库
	_,err = user.AddUser()

	//3、将处理结果返回给客户端

	if err !=nil {
		s.Data["Error"]  = err.Error()
		s.TplName = "404.html"
	}

	s.TplName = "index.html"

}
