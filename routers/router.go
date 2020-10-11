package routers

import (
	"DataCertPlatfom/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/sign", &controllers.SignController{})
    beego.Router("/register",&controllers.RegisterController{})
    beego.Router("/index",&controllers.IndexController{})
}
