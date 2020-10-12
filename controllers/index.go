package controllers

import (
	"DataCertPlatfom/Hash"
	"DataCertPlatfom/db_mysql"
	"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {
	i.TplName = "index.html"
}
func (i *IndexController)Post()  {
	var username = i.Ctx.Request.Form.Get("name")
	var pwd = i.Ctx.Request.Form.Get("pwd")

	fmt.Println(username,pwd)

	pwd = Hash.HASH(pwd,"md5",false)
	fmt.Println(username,pwd)

	falg := db_mysql.Query_user_info(username,pwd)
	if falg == false{
		i.TplName = "404.html"
	}else {
		i.TplName ="register.html"
	}

}