package controllers

import (
	"github.com/astaxie/beego"
)

type NController struct {
	beego.Controller
}

func (p *NController) Get() {
	p.Data["Vápara"] = "Pagina teste"
	p.Data["Email2"] = "survivorplay@gmail.com"
	p.TplName = "teste.html"
}
