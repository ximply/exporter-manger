package config

import (
	"github.com/astaxie/beego"
	"time"
	"strings"
)

type GlobalConfig struct {
	NodeCfg Node
	NginxCfg Nginx
	NginxVtsCfg NginxVts
	PhpfpmCfg Phpfpm
	RedisCfg Redis
	MemcachedCfg Memcached
	MysqlCfg Mysql
	MultiMysqlCfg MultiMysql
	HaproxyCfg Haproxy
	GearmanCfg Gearman
	MongodbCfg Mongodb
	DellHardwareCfg DellHardware
	XenserverCfg Xenserver
	ElasticsearchCfg Elasticsearch
	LogstashCfg Logstash
	MultiLogstashCfg MultiLogstash
	PingCfg Ping
	TcpPingCfg TcpPing
	HttpStatCfg HttpStat
	PingDomainCfg PingDomain
	CertwacherCfg Certwacher
	AliveCfg Alive
	RabbitmqCfg Rabbitmq
	SupervisorCfg Supervisor
	JavaCfg Java
	BeanstalkdCfg Beanstalkd
	BindCfg Bind
	SolrCfg Solr
	HadoopDataNodeCfg HadoopDataNode
	HadoopNameNodeCfg HadoopNameNode
	HadoopSecondNameNodeCfg HadoopSecondNameNode
	HadoopResourceManagerCfg HadoopResourceManager
	KafkaCfg Kafka
	ZookeeperCfg Zookeeper
	CompanyConnCfg CompanyConn
}

type BaseConfig struct {
	Enable bool
	UnixSockFile string
	MetricsPath string
	MetricsRouter string
	Timeout time.Duration
	Filters map[string]string
}

type CompanyConn struct {
	BaseCfg BaseConfig
}

type Node struct {
	BaseCfg BaseConfig
}

type Nginx struct {
	BaseCfg BaseConfig
}

type NginxVts struct {
	BaseCfg BaseConfig
}

type Phpfpm struct {
	BaseCfg BaseConfig
}

type Redis struct {
	BaseCfg BaseConfig
}

type Memcached struct {
	BaseCfg BaseConfig
}

type Mysql struct {
	BaseCfg BaseConfig
}

type MultiMysql struct {
	BaseCfg BaseConfig
}

type Haproxy struct {
	BaseCfg BaseConfig
}

type Gearman struct {
	BaseCfg BaseConfig
}

type Mongodb struct {
	BaseCfg BaseConfig
}

type DellHardware struct {
	BaseCfg BaseConfig
}

type Xenserver struct {
	BaseCfg BaseConfig
}

type Elasticsearch struct {
	BaseCfg BaseConfig
}

type Logstash struct {
	BaseCfg BaseConfig
}

type MultiLogstash struct {
	BaseCfg BaseConfig
}

type Ping struct {
	BaseCfg BaseConfig
}

type TcpPing struct {
	BaseCfg BaseConfig
}

type HttpStat struct {
	BaseCfg BaseConfig
}

type PingDomain struct {
	BaseCfg BaseConfig
}

type Certwacher struct {
	BaseCfg BaseConfig
}

type Alive struct {
	BaseCfg BaseConfig
}

type Rabbitmq struct {
	BaseCfg BaseConfig
}

type Supervisor struct {
	BaseCfg BaseConfig
}

type Java struct {
	BaseCfg BaseConfig
}

type Beanstalkd struct {
	BaseCfg BaseConfig
}

type Bind struct {
	BaseCfg BaseConfig
}

type Solr struct {
	BaseCfg BaseConfig
}

type HadoopDataNode struct {
	BaseCfg BaseConfig
}

type HadoopNameNode struct {
	BaseCfg BaseConfig
}

type HadoopSecondNameNode struct {
	BaseCfg BaseConfig
}

type HadoopResourceManager struct {
	BaseCfg BaseConfig
}

type Kafka struct {
	BaseCfg BaseConfig
}

type Zookeeper struct {
	BaseCfg BaseConfig
}

var globalCfg GlobalConfig

func Init() {
	// company_conn
	globalCfg.CompanyConnCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("company_conn", false)
	if CompanyConnConfig().BaseCfg.Enable {
		globalCfg.CompanyConnCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("company_conn.unix_sock",
			"/dev/shm/companyconn.sock")
		globalCfg.CompanyConnCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("company_conn.metrics_path",
			"/metrics")
		globalCfg.CompanyConnCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("company_conn.metrics_router",
			"/companyconn")
		globalCfg.CompanyConnCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("company_conn.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("company_conn.filters", ""),
			",")
		if len(filters) > 0 {
			var m map[string]string
			m = make(map[string]string)
			for _, s := range filters {
				if len(s) == 0 {
					continue
				}
				m[s] = s
			}
			globalCfg.CompanyConnCfg.BaseCfg.Filters = m
		}
	}

	// node exporter
	globalCfg.NodeCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("node_exporter", false)
	if NodeConfig().BaseCfg.Enable {
		globalCfg.NodeCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("node_exporter.unix_sock",
			"/dev/shm/node_exporter.sock")
		globalCfg.NodeCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("node_exporter.metrics_path",
			"/metrics")
		globalCfg.NodeCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("node_exporter.metrics_router",
			"/node")
		globalCfg.NodeCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("node_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("node_exporter.filters", ""),
			",")
		if len(filters) > 0 {
			var m map[string]string
			m = make(map[string]string)
			for _, s := range filters {
				if len(s) == 0 {
					continue
				}
				m[s] = s
			}
			globalCfg.NodeCfg.BaseCfg.Filters = m
		}
	}

	// nginx exporter
	globalCfg.NginxCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("nginx_exporter", false)
	if NginxConfig().BaseCfg.Enable {
		globalCfg.NginxCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("nginx_exporter.unix_sock",
			"/dev/shm/nginx_exporter.sock")
		globalCfg.NginxCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("nginx_exporter.metrics_path",
			"/metrics")
		globalCfg.NginxCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("nginx_exporter.metrics_router",
			"/nginx")
		globalCfg.NginxCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("nginx_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("nginx_exporter.filters", ""),
			",")
		if len(filters) > 0 {
			var m map[string]string
			m = make(map[string]string)
			for _, s := range filters {
				if len(s) == 0 {
					continue
				}
				m[s] = s
			}
			globalCfg.NginxCfg.BaseCfg.Filters = m
		}
	}

	// nginx vts exporter
	globalCfg.NginxVtsCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("nginxvts_exporter", false)
	if NginxVtsConfig().BaseCfg.Enable {
		globalCfg.NginxVtsCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("nginxvts_exporter.unix_sock",
			"/dev/shm/nginxvts_exporter.sock")
		globalCfg.NginxVtsCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("nginxvts_exporter.metrics_path",
			"/metrics")
		globalCfg.NginxVtsCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("nginxvts_exporter.metrics_router",
			"/nginxvts")
		globalCfg.NginxVtsCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("nginxvts_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("nginxvts_exporter.filters", ""),
			",")
		if len(filters) > 0 {
			var m map[string]string
			m = make(map[string]string)
			for _, s := range filters {
				if len(s) == 0 {
					continue
				}
				m[s] = s
			}
			globalCfg.NginxVtsCfg.BaseCfg.Filters = m
		}
	}

	// php-fpm exporter
	globalCfg.PhpfpmCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("phpfpm_exporter", false)
	if PhpfpmConfig().BaseCfg.Enable {
		globalCfg.PhpfpmCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("phpfpm_exporter.unix_sock",
			"/dev/shm/php-fpm_exporter.sock")
		globalCfg.PhpfpmCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("phpfpm_exporter.metrics_path",
			"/metrics")
		globalCfg.PhpfpmCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("phpfpm_exporter.metrics_router",
			"/phpfpm")
		globalCfg.PhpfpmCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("phpfpm_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("phpfpm_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.PhpfpmCfg.BaseCfg.Filters = m
	}

	// redis exporter
	globalCfg.RedisCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("redis_exporter", false)
	if RedisConfig().BaseCfg.Enable {
		globalCfg.RedisCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("redis_exporter.unix_sock",
			"/dev/shm/redis_exporter.sock")
		globalCfg.RedisCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("redis_exporter.metrics_path",
			"/metrics")
		globalCfg.RedisCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("redis_exporter.metrics_router",
			"/redis")
		globalCfg.RedisCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("redis_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("redis_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.RedisCfg.BaseCfg.Filters = m
	}

	// memcached exporter
	globalCfg.MemcachedCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("memcached_exporter", false)
	if MemcachedConfig().BaseCfg.Enable {
		globalCfg.MemcachedCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("memcached_exporter.unix_sock",
			"/dev/shm/memcached_exporter.sock")
		globalCfg.MemcachedCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("memcached_exporter.metrics_path",
			"/metrics")
		globalCfg.MemcachedCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("memcached_exporter.metrics_router",
			"/memcached")
		globalCfg.MemcachedCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("memcached_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("memcached_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.MemcachedCfg.BaseCfg.Filters = m
	}

	// mysqld exporter
	globalCfg.MysqlCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("mysqld_exporter", false)
	if MysqlConfig().BaseCfg.Enable {
		globalCfg.MysqlCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("mysqld_exporter.unix_sock",
			"/dev/shm/mysqld_exporter.sock")
		globalCfg.MysqlCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("mysqld_exporter.metrics_path",
			"/metrics")
		globalCfg.MysqlCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("mysqld_exporter.metrics_router",
			"/mysql")
		globalCfg.MysqlCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("mysqld_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("mysqld_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.MysqlCfg.BaseCfg.Filters = m
	}

	// multi mysqld exporter
	globalCfg.MultiMysqlCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("multimysqld_exporter", false)
	if MultiMysqlConfig().BaseCfg.Enable {
		globalCfg.MultiMysqlCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("multimysqld_exporter.unix_sock",
			"/dev/shm/multimysqld_exporter.sock")
		globalCfg.MultiMysqlCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("multimysqld_exporter.metrics_path",
			"/metrics")
		globalCfg.MultiMysqlCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("multimysqld_exporter.metrics_router",
			"/mysqls")
		globalCfg.MultiMysqlCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("multimysqld_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("multimysqld_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.MultiMysqlCfg.BaseCfg.Filters = m
	}

	// haproxy exporter
	globalCfg.HaproxyCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("haproxy_exporter", false)
	if HaproxyConfig().BaseCfg.Enable {
		globalCfg.HaproxyCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("haproxy_exporter.unix_sock",
			"/dev/shm/haproxy_exporter.sock")
		globalCfg.HaproxyCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("haproxy_exporter.metrics_path",
			"/metrics")
		globalCfg.HaproxyCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("haproxy_exporter.metrics_router",
			"/haproxy")
		globalCfg.HaproxyCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("haproxy_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("haproxy_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.HaproxyCfg.BaseCfg.Filters = m
	}

	// gearman exporter
	globalCfg.GearmanCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("gearman_exporter", false)
	if GearmanConfig().BaseCfg.Enable {
		globalCfg.GearmanCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("gearman_exporter.unix_sock",
			"/dev/shm/gearman_exporter.sock")
		globalCfg.GearmanCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("gearman_exporter.metrics_path",
			"/metrics")
		globalCfg.GearmanCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("gearman_exporter.metrics_router",
			"/gearman")
		globalCfg.GearmanCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("gearman_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("gearman_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.GearmanCfg.BaseCfg.Filters = m
	}

	// mongodb exporter
	globalCfg.MongodbCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("mongodb_exporter", false)
	if MongodbConfig().BaseCfg.Enable {
		globalCfg.MongodbCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("mongodb_exporter.unix_sock",
			"/dev/shm/mongodb_exporter.sock")
		globalCfg.MongodbCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("mongodb_exporter.metrics_path",
			"/metrics")
		globalCfg.MongodbCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("mongodb_exporter.metrics_router",
			"/mongodb")
		globalCfg.MongodbCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("mongodb_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("mongodb_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.MongodbCfg.BaseCfg.Filters = m
	}

	// dell hardware exporter
	globalCfg.DellHardwareCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("dellhardware_exporter", false)
	if DellHardwareConfig().BaseCfg.Enable {
		globalCfg.DellHardwareCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("dellhardware_exporter.unix_sock",
			"/dev/shm/dellhardware_exporter.sock")
		globalCfg.DellHardwareCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("dellhardware_exporter.metrics_path",
			"/metrics")
		globalCfg.DellHardwareCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("dellhardware_exporter.metrics_router",
			"/dellhw")
		globalCfg.DellHardwareCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("dellhardware_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("dellhardware_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.DellHardwareCfg.BaseCfg.Filters = m
	}

	// xenserver exporter
	globalCfg.XenserverCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("xenserver_exporter", false)
	if XenserverConfig().BaseCfg.Enable {
		globalCfg.XenserverCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("xenserver_exporter.unix_sock",
			"/dev/shm/xenserver_exporter.sock")
		globalCfg.XenserverCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("xenserver_exporter.metrics_path",
			"/metrics")
		globalCfg.XenserverCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("xenserver_exporter.metrics_router",
			"/xenserver")
		globalCfg.XenserverCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("xenserver_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("xenserver_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.XenserverCfg.BaseCfg.Filters = m
	}

	// elasticsearch exporter
	globalCfg.ElasticsearchCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("elasticsearch_exporter", false)
	if ElasticsearchConfig().BaseCfg.Enable {
		globalCfg.ElasticsearchCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("elasticsearch_exporter.unix_sock",
			"/dev/shm/elasticsearch_exporter.sock")
		globalCfg.ElasticsearchCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("elasticsearch_exporter.metrics_path",
			"/metrics")
		globalCfg.ElasticsearchCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("elasticsearch_exporter.metrics_router",
			"/es")
		globalCfg.ElasticsearchCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("elasticsearch_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("elasticsearch_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.ElasticsearchCfg.BaseCfg.Filters = m
	}

	// logstash exporter
	globalCfg.LogstashCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("logstash_exporter", false)
	if LogstashConfig().BaseCfg.Enable {
		globalCfg.LogstashCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("logstash_exporter.unix_sock",
			"/dev/shm/logstash_exporter.sock")
		globalCfg.LogstashCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("logstash_exporter.metrics_path",
			"/metrics")
		globalCfg.LogstashCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("logstash_exporter.metrics_router",
			"/logstash")
		globalCfg.LogstashCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("logstash_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("logstash_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.LogstashCfg.BaseCfg.Filters = m
	}

	// multi logstash exporter
	globalCfg.MultiLogstashCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("multilogstash_exporter", false)
	if MultiLogstashConfig().BaseCfg.Enable {
		globalCfg.MultiLogstashCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("multilogstash_exporter.unix_sock",
			"/dev/shm/multilogstash_exporter.sock")
		globalCfg.MultiLogstashCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("multilogstash_exporter.metrics_path",
			"/metrics")
		globalCfg.MultiLogstashCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("multilogstash_exporter.metrics_router",
			"/logstashes")
		globalCfg.MultiLogstashCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("multilogstash_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("multilogstash_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.MultiLogstashCfg.BaseCfg.Filters = m
	}

	// ping exporter
	globalCfg.PingCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("ping_exporter", false)
	if PingConfig().BaseCfg.Enable {
		globalCfg.PingCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("ping_exporter.unix_sock",
			"/dev/shm/ping_exporter.sock")
		globalCfg.PingCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("ping_exporter.metrics_path",
			"/metrics")
		globalCfg.PingCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("ping_exporter.metrics_router",
			"/ping")
		globalCfg.PingCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("ping_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("ping_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.PingCfg.BaseCfg.Filters = m
	}

	// tcp ping exporter
	globalCfg.TcpPingCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("tcpping_exporter", false)
	if TcpPingConfig().BaseCfg.Enable {
		globalCfg.TcpPingCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("tcpping_exporter.unix_sock",
			"/dev/shm/tcpping_exporter.sock")
		globalCfg.TcpPingCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("tcpping_exporter.metrics_path",
			"/metrics")
		globalCfg.TcpPingCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("tcpping_exporter.metrics_router",
			"/tcpping")
		globalCfg.TcpPingCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("tcpping_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("tcpping_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.TcpPingCfg.BaseCfg.Filters = m
	}

	// httpstat exporter
	globalCfg.HttpStatCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("httpstat_exporter", false)
	if HttpStatConfig().BaseCfg.Enable {
		globalCfg.HttpStatCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("httpstat_exporter.unix_sock",
			"/dev/shm/httpstat_exporter.sock")
		globalCfg.HttpStatCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("httpstat_exporter.metrics_path",
			"/metrics")
		globalCfg.HttpStatCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("httpstat_exporter.metrics_router",
			"/httpstat")
		globalCfg.HttpStatCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("httpstat_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("httpstat_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.HttpStatCfg.BaseCfg.Filters = m
	}

	// ping domain exporter
	globalCfg.PingDomainCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("pingdomain_exporter", false)
	if PingDomainConfig().BaseCfg.Enable {
		globalCfg.PingDomainCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("pingdomain_exporter.unix_sock",
			"/dev/shm/pingdomain_exporter.sock")
		globalCfg.PingDomainCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("pingdomain_exporter.metrics_path",
			"/metrics")
		globalCfg.PingDomainCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("pingdomain_exporter.metrics_router",
			"/pingdomain")
		globalCfg.PingDomainCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("pingdomain_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("pingdomain_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.PingDomainCfg.BaseCfg.Filters = m
	}

	// certwacher exporter
	globalCfg.CertwacherCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("cert_exporter", false)
	if CertwacherConfig().BaseCfg.Enable {
		globalCfg.CertwacherCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("cert_exporter.unix_sock",
			"/dev/shm/cert_exporter.sock")
		globalCfg.CertwacherCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("cert_exporter.metrics_path",
			"/metrics")
		globalCfg.CertwacherCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("cert_exporter.metrics_router",
			"/cert")
		globalCfg.CertwacherCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("cert_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("cert_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.CertwacherCfg.BaseCfg.Filters = m
	}

	// alive exporter
	globalCfg.AliveCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("alive_exporter", false)
	if AliveConfig().BaseCfg.Enable {
		globalCfg.AliveCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("alive_exporter.unix_sock",
			"/dev/shm/alive_exporter.sock")
		globalCfg.AliveCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("alive_exporter.metrics_path",
			"/metrics")
		globalCfg.AliveCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("alive_exporter.metrics_router",
			"/alive")
		globalCfg.AliveCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("alive_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("alive_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.AliveCfg.BaseCfg.Filters = m
	}

	// rabbitmq exporter
	globalCfg.RabbitmqCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("rabbitmq_exporter", false)
	if RabbitmqConfig().BaseCfg.Enable {
		globalCfg.RabbitmqCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("rabbitmq_exporter.unix_sock",
			"/dev/shm/rabbitmq_exporter.sock")
		globalCfg.RabbitmqCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("rabbitmq_exporter.metrics_path",
			"/metrics")
		globalCfg.RabbitmqCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("rabbitmq_exporter.metrics_router",
			"/rabbitmq")
		globalCfg.RabbitmqCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("rabbitmq_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("rabbitmq_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.RabbitmqCfg.BaseCfg.Filters = m
	}

	// supervisor exporter
	globalCfg.SupervisorCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("supervisor_exporter", false)
	if SupervisorConfig().BaseCfg.Enable {
		globalCfg.SupervisorCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("supervisor_exporter.unix_sock",
			"/dev/shm/supervisor_exporter.sock")
		globalCfg.SupervisorCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("supervisor_exporter.metrics_path",
			"/metrics")
		globalCfg.SupervisorCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("supervisor_exporter.metrics_router",
			"/rabbitmq")
		globalCfg.SupervisorCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("supervisor_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("supervisor_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.SupervisorCfg.BaseCfg.Filters = m
	}

	// java exporter
	globalCfg.JavaCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("java_exporter", false)
	if JavaConfig().BaseCfg.Enable {
		globalCfg.JavaCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("java_exporter.metrics_router",
			"/java")
		globalCfg.JavaCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("java_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("java_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.JavaCfg.BaseCfg.Filters = m
	}

	// beanstalkd exporter
	globalCfg.BeanstalkdCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("beanstalkd_exporter", false)
	if BeanstalkdConfig().BaseCfg.Enable {
		globalCfg.BeanstalkdCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("beanstalkd_exporter.unix_sock",
			"/dev/shm/beanstalkd_exporter.sock")
		globalCfg.BeanstalkdCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("beanstalkd_exporter.metrics_path",
			"/metrics")
		globalCfg.BeanstalkdCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("beanstalkd_exporter.metrics_router",
			"/beanstalkd")
		globalCfg.BeanstalkdCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("beanstalkd_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("beanstalkd_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.BeanstalkdCfg.BaseCfg.Filters = m
	}

	// bind exporter
	globalCfg.BindCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("bind_exporter", false)
	if BindConfig().BaseCfg.Enable {
		globalCfg.BindCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("bind_exporter.unix_sock",
			"/dev/shm/bind_exporter.sock")
		globalCfg.BindCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("bind_exporter.metrics_path",
			"/metrics")
		globalCfg.BindCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("bind_exporter.metrics_router",
			"/bind")
		globalCfg.BindCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("bind_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("bind_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.BindCfg.BaseCfg.Filters = m
	}
	
	// solr exporter
	globalCfg.SolrCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("solr_exporter", false)
	if SolrConfig().BaseCfg.Enable {
		globalCfg.SolrCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("solr_exporter.unix_sock",
			"/dev/shm/solr_exporter.sock")
		globalCfg.SolrCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("solr_exporter.metrics_path",
			"/metrics")
		globalCfg.SolrCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("solr_exporter.metrics_router",
			"/bind")
		globalCfg.SolrCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("solr_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("solr_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.SolrCfg.BaseCfg.Filters = m
	}

	// hadoop data node exporter
	globalCfg.HadoopDataNodeCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("hadoop_datanode_exporter", false)
	if HadoopDataNodeConfig().BaseCfg.Enable {
		globalCfg.HadoopDataNodeCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("hadoop_datanode_exporter.unix_sock",
			"/dev/shm/hadoop_datanode_exporter.sock")
		globalCfg.HadoopDataNodeCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("hadoop_datanode_exporter.metrics_path",
			"/metrics")
		globalCfg.HadoopDataNodeCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("hadoop_datanode_exporter.metrics_router",
			"/hadoopdn")
		globalCfg.HadoopDataNodeCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("hadoop_datanode_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("hadoop_datanode_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.HadoopDataNodeCfg.BaseCfg.Filters = m
	}

	// hadoop name node exporter
	globalCfg.HadoopNameNodeCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("hadoop_namenode_exporter", false)
	if HadoopNameNodeConfig().BaseCfg.Enable {
		globalCfg.HadoopNameNodeCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("hadoop_namenode_exporter.unix_sock",
			"/dev/shm/hadoop_namenode_exporter.sock")
		globalCfg.HadoopNameNodeCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("hadoop_namenode_exporter.metrics_path",
			"/metrics")
		globalCfg.HadoopNameNodeCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("hadoop_namenode_exporter.metrics_router",
			"/hadoopnn")
		globalCfg.HadoopNameNodeCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("hadoop_namenode_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("hadoop_namenode_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.HadoopNameNodeCfg.BaseCfg.Filters = m
	}

	// hadoop second name node exporter
	globalCfg.HadoopSecondNameNodeCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("hadoop_secondnamenode_exporter", false)
	if HadoopSecondNameNodeConfig().BaseCfg.Enable {
		globalCfg.HadoopSecondNameNodeCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("hadoop_secondnamenode_exporter.unix_sock",
			"/dev/shm/hadoop_secondnamenode_exporter.sock")
		globalCfg.HadoopSecondNameNodeCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("hadoop_secondnamenode_exporter.metrics_path",
			"/metrics")
		globalCfg.HadoopSecondNameNodeCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("hadoop_secondnamenode_exporter.metrics_router",
			"/hadoopsnn")
		globalCfg.HadoopSecondNameNodeCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("hadoop_secondnamenode_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("hadoop_secondnamenode_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.HadoopSecondNameNodeCfg.BaseCfg.Filters = m
	}

	// hadoop resource manager exporter
	globalCfg.HadoopResourceManagerCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("hadoop_resourcemanager_exporter", false)
	if HadoopResourceManagerConfig().BaseCfg.Enable {
		globalCfg.HadoopResourceManagerCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("hadoop_resourcemanager_exporter.unix_sock",
			"/dev/shm/hadoop_resourcemanager_exporter.sock")
		globalCfg.HadoopResourceManagerCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("hadoop_resourcemanager_exporter.metrics_path",
			"/metrics")
		globalCfg.HadoopResourceManagerCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("hadoop_resourcemanager_exporter.metrics_router",
			"/hadooprm")
		globalCfg.HadoopResourceManagerCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("hadoop_resourcemanager_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("hadoop_resourcemanager_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.HadoopResourceManagerCfg.BaseCfg.Filters = m
	}

	// kafka exporter
	globalCfg.KafkaCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("kafka_exporter", false)
	if KafkaConfig().BaseCfg.Enable {
		globalCfg.KafkaCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("kafka_exporter.unix_sock",
			"/dev/shm/kafka_exporter.sock")
		globalCfg.KafkaCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("kafka_exporter.metrics_path",
			"/metrics")
		globalCfg.KafkaCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("kafka_exporter.metrics_router",
			"/kafka")
		globalCfg.KafkaCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("kafka_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("kafka_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.KafkaCfg.BaseCfg.Filters = m
	}

	// zookeeper exporter
	globalCfg.ZookeeperCfg.BaseCfg.Enable = beego.AppConfig.DefaultBool("zookeeper_exporter", false)
	if ZookeeperConfig().BaseCfg.Enable {
		globalCfg.ZookeeperCfg.BaseCfg.UnixSockFile = beego.AppConfig.DefaultString("zookeeper_exporter.unix_sock",
			"/dev/shm/kafka_exporter.sock")
		globalCfg.ZookeeperCfg.BaseCfg.MetricsPath = beego.AppConfig.DefaultString("zookeeper_exporter.metrics_path",
			"/metrics")
		globalCfg.ZookeeperCfg.BaseCfg.MetricsRouter = beego.AppConfig.DefaultString("zookeeper_exporter.metrics_router",
			"/zk")
		globalCfg.ZookeeperCfg.BaseCfg.Timeout = time.Duration(beego.AppConfig.DefaultInt("zookeeper_exporter.timeout",
			5)) * time.Second
		filters := strings.Split(beego.AppConfig.DefaultString("zookeeper_exporter.filters", ""),
			",")
		var m map[string]string
		m = make(map[string]string)
		for _, s := range filters {
			if len(s) == 0 {
				continue
			}
			m[s] = s
		}
		globalCfg.ZookeeperCfg.BaseCfg.Filters = m
	}
}

func CompanyConnConfig() CompanyConn {
	return globalCfg.CompanyConnCfg
}

func NodeConfig() Node {
	return globalCfg.NodeCfg
}

func NginxConfig() Nginx {
	return globalCfg.NginxCfg
}

func NginxVtsConfig() NginxVts {
	return globalCfg.NginxVtsCfg
}

func PhpfpmConfig() Phpfpm {
	return globalCfg.PhpfpmCfg
}

func RedisConfig() Redis {
	return globalCfg.RedisCfg
}

func MemcachedConfig() Memcached {
	return globalCfg.MemcachedCfg
}

func MysqlConfig() Mysql {
	return globalCfg.MysqlCfg
}

func MultiMysqlConfig() MultiMysql {
	return globalCfg.MultiMysqlCfg
}

func HaproxyConfig() Haproxy {
	return globalCfg.HaproxyCfg
}

func GearmanConfig() Gearman {
	return globalCfg.GearmanCfg
}

func MongodbConfig() Mongodb {
	return globalCfg.MongodbCfg
}

func DellHardwareConfig() DellHardware {
	return globalCfg.DellHardwareCfg
}

func XenserverConfig() Xenserver {
	return globalCfg.XenserverCfg
}

func ElasticsearchConfig() Elasticsearch {
	return globalCfg.ElasticsearchCfg
}

func LogstashConfig() Logstash {
	return globalCfg.LogstashCfg
}

func MultiLogstashConfig() MultiLogstash {
	return globalCfg.MultiLogstashCfg
}

func PingConfig() Ping {
	return globalCfg.PingCfg
}

func TcpPingConfig() TcpPing {
	return globalCfg.TcpPingCfg
}

func HttpStatConfig() HttpStat {
	return globalCfg.HttpStatCfg
}

func PingDomainConfig() PingDomain {
	return globalCfg.PingDomainCfg
}

func CertwacherConfig() Certwacher {
	return globalCfg.CertwacherCfg
}

func AliveConfig() Alive {
	return globalCfg.AliveCfg
}

func RabbitmqConfig() Rabbitmq {
	return globalCfg.RabbitmqCfg
}

func SupervisorConfig() Supervisor {
	return globalCfg.SupervisorCfg
}

func JavaConfig() Java {
	return globalCfg.JavaCfg
}

func BeanstalkdConfig() Beanstalkd {
	return globalCfg.BeanstalkdCfg
}

func BindConfig() Bind {
	return globalCfg.BindCfg
}

func SolrConfig() Solr {
	return globalCfg.SolrCfg
}

func HadoopDataNodeConfig() HadoopDataNode {
	return globalCfg.HadoopDataNodeCfg
}

func HadoopNameNodeConfig() HadoopNameNode {
	return globalCfg.HadoopNameNodeCfg
}

func HadoopSecondNameNodeConfig() HadoopSecondNameNode {
	return globalCfg.HadoopSecondNameNodeCfg
}

func HadoopResourceManagerConfig() HadoopResourceManager {
	return globalCfg.HadoopResourceManagerCfg
}

func KafkaConfig() Kafka {
	return globalCfg.KafkaCfg
}

func ZookeeperConfig() Zookeeper {
	return globalCfg.ZookeeperCfg
}