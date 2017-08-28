package routers

import (
	"ProjetoWeb/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	beego.Router("/teste", &controllers.NovoController{}, "get:Testa")
}
