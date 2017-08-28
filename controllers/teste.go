package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type NovoController struct {
	beego.Controller
}

func (n *NovoController) Get() {
	n.Data["Nome"] = "Thomas"
	n.Data["Idade"] = "21"
	n.TplName = "teste.html"
}

func (n *NovoController) Testa() {
	n.TplName = "teste.html"
	fmt.Println("oi")
}
