package main

import (
	"DataCertPlatfom/db_mysql"
	_ "DataCertPlatfom/routers"
	"github.com/astaxie/beego"
)

func main() {
	db_mysql.OpenDB()
	//静态资源文件映射
	//				url 出现 img时			去path路径里 ./static/xxx 路径去找相应的文件
	beego.SetStaticPath("../static/baoqaun", "./static/baoquan")//static
	//beego.SetStaticPath("/css", "./static/css")//static
	//beego.SetStaticPath("/js", "./static/js")//static
	beego.Run()
}

