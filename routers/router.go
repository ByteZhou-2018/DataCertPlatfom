package routers

import (
	"DataCertPlatfom/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/sign_in", &controllers.SignController{})
    beego.Router("/sign_in.html", &controllers.SignController{})

    beego.Router("login",&controllers.LoginController{})
    beego.Router("/login.html",&controllers.LoginController{})

    beego.Router("/files",&controllers.FilesController{})
    beego.Router("/files.html",&controllers.FilesController{})

    beego.Router("/home.html",&controllers.HomeController{})
    beego.Router("/home",&controllers.HomeController{})

}
