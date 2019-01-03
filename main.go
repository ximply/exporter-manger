package main

import (
	"github.com/ximply/exporter-manger/config"
	_ "github.com/ximply/exporter-manger/routers"
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/models"
	"github.com/ximply/exporter-manger/uploads"
)

func main() {
	models.Init()
	if config.UConfigs().On {
		uploads.Start()
	}
	beego.Run()
}
