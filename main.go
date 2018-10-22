package main

import (
	_ "github.com/ximply/exporter-manger/routers"
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/models"
	"github.com/ximply/exporter-manger/uploads"
)

func main() {
	models.Init()
	uploads.Start()
	beego.Run()
}
