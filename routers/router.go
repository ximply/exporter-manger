package routers

import (
	"github.com/ximply/exporter-manger/controllers"
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/config"
)

func init() {
    config.Init()
    beego.Router("/", &controllers.MainController{})
    // node exporter
	if config.NodeConfig().BaseCfg.Enable {
		beego.Router(config.NodeConfig().BaseCfg.MetricsRouter,
			&controllers.NodeController{}, "get:NodeMetrics")
	}
	// redis exporter
	if config.RedisConfig().BaseCfg.Enable {
		beego.Router(config.RedisConfig().BaseCfg.MetricsRouter,
			&controllers.RedisController{}, "get:RedisMetrics")
	}
	// memcached exporter
	if config.MemcachedConfig().BaseCfg.Enable {
		beego.Router(config.MemcachedConfig().BaseCfg.MetricsRouter,
			&controllers.MemchachedController{}, "get:MemcachedMetrics")
	}
}
