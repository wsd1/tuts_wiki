package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["Website"] = "d1works.com"
	c.Data["Email"] = "yizuoshe@gmail.com"
	c.Data["Version"] = "0.1"
	c.Data["IsLogin"] = true

	c.TplNames = "login.html"
}
