package routers

import (
	"github.com/ximply/exporter-manger/controllers"
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/config"
)

func init() {
    config.Init()
    beego.Router("/", &controllers.MainController{})
	if config.NodeConfig().BaseCfg.Enable {
		beego.Router(config.NodeConfig().BaseCfg.MetricsRouter,
			&controllers.NodeController{}, "get:NodeMetrics")
	}
}
