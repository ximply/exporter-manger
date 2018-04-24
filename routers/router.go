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
	// nginx exporter
	if config.NginxConfig().BaseCfg.Enable {
		beego.Router(config.NginxConfig().BaseCfg.MetricsRouter,
			&controllers.NginxController{}, "get:NginxMetrics")
	}
	// nginx vts exporter
	if config.NginxVtsConfig().BaseCfg.Enable {
		beego.Router(config.NginxVtsConfig().BaseCfg.MetricsRouter,
			&controllers.NginxVtsController{}, "get:NginxVtsMetrics")
	}
	// php-fpm exporter
	if config.PhpfpmConfig().BaseCfg.Enable {
		beego.Router(config.PhpfpmConfig().BaseCfg.MetricsRouter,
			&controllers.PhpfpmController{}, "get:PhpfpmMetrics")
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
	// mysqld exporter
	if config.MysqlConfig().BaseCfg.Enable {
		beego.Router(config.MysqlConfig().BaseCfg.MetricsRouter,
			&controllers.MysqldController{}, "get:MysqldMetrics")
	}
	// haproxy exporter
	if config.HaproxyConfig().BaseCfg.Enable {
		beego.Router(config.HaproxyConfig().BaseCfg.MetricsRouter,
			&controllers.HaproxyController{}, "get:HaproxyMetrics")
	}
}
