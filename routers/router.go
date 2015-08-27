package routers

import (
	"github.com/astaxie/beego"
	"tuts_wiki/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/home", &controllers.HomeController{})
}
