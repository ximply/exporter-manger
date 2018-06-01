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
	// multi mysqld exporter
	if config.MultiMysqlConfig().BaseCfg.Enable {
		beego.Router(config.MultiMysqlConfig().BaseCfg.MetricsRouter,
			&controllers.MultiMysqldController{}, "get:MultiMysqldMetrics")
	}
	// haproxy exporter
	if config.HaproxyConfig().BaseCfg.Enable {
		beego.Router(config.HaproxyConfig().BaseCfg.MetricsRouter,
			&controllers.HaproxyController{}, "get:HaproxyMetrics")
	}
	// gearman exporter
	if config.GearmanConfig().BaseCfg.Enable {
		beego.Router(config.GearmanConfig().BaseCfg.MetricsRouter,
			&controllers.GearmanController{}, "get:GearmanMetrics")
	}
	// mongodb exporter
	if config.MongodbConfig().BaseCfg.Enable {
		beego.Router(config.MongodbConfig().BaseCfg.MetricsRouter,
			&controllers.MongodbController{}, "get:MongodbMetrics")
	}
	// dellhardware exporter
	if config.DellHardwareConfig().BaseCfg.Enable {
		beego.Router(config.DellHardwareConfig().BaseCfg.MetricsRouter,
			&controllers.DellHardwareController{}, "get:DellHardwareMetrics")
	}

	// xenserver exporter
	if config.XenserverConfig().BaseCfg.Enable {
		beego.Router(config.XenserverConfig().BaseCfg.MetricsRouter,
			&controllers.XenserverController{}, "get:XenserverMetrics")
	}

	// elasticsearch exporter
	if config.ElasticsearchConfig().BaseCfg.Enable {
		beego.Router(config.ElasticsearchConfig().BaseCfg.MetricsRouter,
			&controllers.ElasticsearchController{}, "get:ElasticsearchMetrics")
	}

	// logstash exporter
	if config.LogstashConfig().BaseCfg.Enable {
		beego.Router(config.LogstashConfig().BaseCfg.MetricsRouter,
			&controllers.LogstashController{}, "get:LogstashMetrics")
	}

	// multi logstash exporter
	if config.MultiLogstashConfig().BaseCfg.Enable {
		beego.Router(config.MultiLogstashConfig().BaseCfg.MetricsRouter,
			&controllers.MultiLogstashController{}, "get:MultiLogstashMetrics")
	}

	// ping exporter
	if config.PingConfig().BaseCfg.Enable {
		beego.Router(config.PingConfig().BaseCfg.MetricsRouter,
			&controllers.PingController{}, "get:PingMetrics")
	}

	// tcp ping exporter
	if config.TcpPingConfig().BaseCfg.Enable {
		beego.Router(config.TcpPingConfig().BaseCfg.MetricsRouter,
			&controllers.TcpPingController{}, "get:TcpPingMetrics")
	}

	// httpstat exporter
	if config.HttpStatConfig().BaseCfg.Enable {
		beego.Router(config.HttpStatConfig().BaseCfg.MetricsRouter,
			&controllers.HttpStatController{}, "get:HttpStatMetrics")
	}

	// ping domain exporter
	if config.PingDomainConfig().BaseCfg.Enable {
		beego.Router(config.PingDomainConfig().BaseCfg.MetricsRouter,
			&controllers.PingDomainController{}, "get:PingDomainMetrics")
	}

	// certwacher exporter
	if config.CertwacherConfig().BaseCfg.Enable {
		beego.Router(config.CertwacherConfig().BaseCfg.MetricsRouter,
			&controllers.CertwacherController{}, "get:CertwacherMetrics")
	}

	// alive exporter
	if config.AliveConfig().BaseCfg.Enable {
		beego.Router(config.AliveConfig().BaseCfg.MetricsRouter,
			&controllers.AliveController{}, "get:AliveMetrics")
	}

	// rabbitmq exporter
	if config.RabbitmqConfig().BaseCfg.Enable {
		beego.Router(config.RabbitmqConfig().BaseCfg.MetricsRouter,
			&controllers.RabbitmqController{}, "get:RabbitmqMetrics")
	}

	// supervisor exporter
	if config.SupervisorConfig().BaseCfg.Enable {
		beego.Router(config.SupervisorConfig().BaseCfg.MetricsRouter,
			&controllers.SupervisorController{}, "get:SupervisorMetrics")
	}

	// java exporter
	if config.JavaConfig().BaseCfg.Enable {
		beego.Router(config.JavaConfig().BaseCfg.MetricsRouter,
			&controllers.JavaController{}, "get:JavaMetrics")
	}
	beego.Router("/javainfo", &controllers.JavaInfoController{},"post:JavaInfo")

	// beanstalkd exporter
	if config.BeanstalkdConfig().BaseCfg.Enable {
		beego.Router(config.BeanstalkdConfig().BaseCfg.MetricsRouter,
			&controllers.BeanstalkdController{}, "get:BeanstalkdMetrics")
	}

	// bind exporter
	if config.BindConfig().BaseCfg.Enable {
		beego.Router(config.BindConfig().BaseCfg.MetricsRouter,
			&controllers.BindController{}, "get:BindMetrics")
	}
}
