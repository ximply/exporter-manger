package main

import (
	_ "github.com/ximply/exporter-manger/routers"
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/models"
)

func main() {
	models.Init()
	beego.Run()
}
