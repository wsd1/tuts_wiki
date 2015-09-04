package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	//	"log"
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

	isExit := c.Input().Get("exit") == "true"

	if isExit {
		c.Ctx.SetCookie("usremail", "", -1, "/words")
		c.Ctx.SetCookie("usrpwd", "", -1, "/words")
		c.Redirect("/", 302)
	}

}

func (c *LoginController) Post() {

	auth_usr := beego.AppConfig.String("usremail")
	if auth_usr == "" {
		auth_usr = "a@b.c"

	}
	auth_pwd := beego.AppConfig.String("usrpwd")
	if auth_pwd == "" {
		auth_pwd = "abc"
	}

	email := c.Input().Get("usremail")
	pwd := c.Input().Get("usrpwd")

	/*
		log.Println("auth_mail:", auth_usr)
		log.Println("auth_pwd:", auth_usr)
		log.Println("in_mail:", email)
		log.Println("in_pwd:", pwd)
	*/

	if auth_usr == email && auth_pwd == pwd {
		maxAge := 0
		if c.Input().Get("usrautologin") == "on" {
			maxAge = 3600
		}

		c.Ctx.SetCookie("usremail", email, maxAge, "/words")
		c.Ctx.SetCookie("usrpwd", pwd, maxAge, "/words")
	}

	c.Redirect("/words", 302)

}

func isLogin(ctx *context.Context) bool {

	auth_usr := beego.AppConfig.String("usremail")
	if auth_usr == "" {
		auth_usr = "a@b.c"

	}
	auth_pwd := beego.AppConfig.String("usrpwd")
	if auth_pwd == "" {
		auth_pwd = "abc"
	}

	ck, err := ctx.Request.Cookie("usremail")
	if err != nil {
		return false
	}
	email := ck.Value

	ck, err = ctx.Request.Cookie("usrpwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	/*
		log.Println("auth_mail:", auth_usr)
		log.Println("auth_pwd:", auth_usr)
		log.Println("in_mail:", email)
		log.Println("in_pwd:", pwd)
	*/
	return auth_usr == email && auth_pwd == pwd
}
