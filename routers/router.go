package routers

import (
	"ProjetoWeb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Router("/", &controllers.MainController{})
}
