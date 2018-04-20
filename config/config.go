package config

import (
	"github.com/astaxie/beego"
	"time"
	"strings"
)

type GlobalConfig struct {
	NodeCfg Node
	RedisCfg Redis
}

type BaseConfig struct {
	Enable bool
	UnixSockFile string
	MetricsPath string
	MetricsRouter string
	Timeout time.Duration
	Filters map[string]string
}

type Node struct {
	BaseCfg BaseConfig
}

type Redis struct {
	BaseCfg BaseConfig
}

var globalCfg GlobalConfig

func Init() {
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
}

func NodeConfig() Node {
	return globalCfg.NodeCfg
}

func RedisConfig() Redis {
	return globalCfg.RedisCfg
}