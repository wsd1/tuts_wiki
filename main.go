package main

import (
	"github.com/astaxie/beego"
	_ "tuts_wiki/routers"
)

func main() {

	beego.SessionOn = true

	beego.Run()
}
