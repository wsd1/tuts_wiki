package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "ExMem"
	c.Data["Email"] = "yizuoshe@gmail.com"
	c.TplNames = "index.html"
}
