package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	v := c.GetSession("asta")
	if v == nil {
		c.SetSession("asta", int(1))
		c.Data["num"] = 0
	} else {
		c.SetSession("asta", v.(int)+1)
		c.Data["num"] = v.(int)
	}

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}