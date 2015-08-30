package main

import (
	"github.com/astaxie/beego"
	"tuts_wiki/models"
	_ "tuts_wiki/routers"
)

func main() {

	models.Init()

	beego.SessionOn = true
	beego.Run()
}
