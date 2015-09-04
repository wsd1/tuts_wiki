package main

import (
	"github.com/astaxie/beego"
	"time"
	"tuts_wiki/models"
	_ "tuts_wiki/routers"
)

func utc2str(in float64) (out string) {
	return time.Unix(int64(in), 0).Format("060102 03:04:05 PM")
}

func main() {

	models.Init()

	beego.AddFuncMap("utc2str", utc2str)

	beego.SessionOn = true
	beego.CopyRequestBody = true

	beego.Run()
}
