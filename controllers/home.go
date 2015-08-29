package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["Website"] = "d1works.com"
	c.Data["Email"] = "yizuoshe@gmail.com"
	c.Data["Version"] = "0.1"
	c.TplNames = "home.html"

	if !isLogin(c.Ctx) {
		c.Redirect("/", 302)
	}

	// view context, init if not defined
	history := c.GetSession("WordPath")
	current := c.GetSession("WordNow")
	if history == nil {
		default_start := beego.AppConfig.String("StartPoint")
		history = []string{default_start}
		c.SetSession("WordPath", history)
	}
	if current == nil {
		current = history.([]string)[0]
		c.SetSession("WordNow", current)
	}

	//select word
	WordSelect := c.Input().Get("select")
	if WordSelect != "" {

		if !isIn(WordSelect, history.([]string)) {
			history = append(history.([]string), WordSelect)
		}

		current = WordSelect
		c.SetSession("WordPath", history)
		c.SetSession("WordNow", current)
	}

	c.Data["WordPath"] = history
	c.Data["WordCurrent"] = current

}

func isIn(str string, strs []string) bool {
	for _, v := range strs {
		if v == str {
			return true
		}
	}
	return false
}
