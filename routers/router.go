package routers

import (
	"DataCertPlatfom/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/sign_in", &controllers.SignController{})
    beego.Router("/sign_in.html", &controllers.SignController{})
    beego.Router("index",&controllers.LoginController{})
    beego.Router("/index.html",&controllers.LoginController{})
}
