package controllers

import (
	"DataCertPlatfom/db_mysql"
	"DataCertPlatfom/models"
	"fmt"

	//"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}
//   处理 register 页面的 post请求
func (r RegisterController)Post()  {
	//1、解析用户页面提交的请求数据
	var user models.User
	user.Phone = r.Ctx.Request.Form.Get("phone")
	user.Password =r.Ctx.Request.Form.Get("password")
	fmt.Println(user)
	//err := r.ParseForm(&user)
	//if err != nil{
	//	 r.Ctx.WriteString("数据解析失败，请重试！")
	//}
	//fmt.Println(user)
	//2、将解析到的数保存到数据库
	_,err := db_mysql.Inseret(user)

	//3、将处理结果返回给客户端

	//3.2 失败。。。。 则提示错误信息

	if err !=nil {
		//r.TplName = "register.html"
		//r.Ctx.WriteString("注册失败，请重试！")

		r.TplName = "404.html"
	}
	//3.1  成功 。。。 跳转到登录界面
	r.TplName = "index.html"

}