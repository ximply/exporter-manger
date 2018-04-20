package models

import (
	"github.com/ximply/exporter-manger/utils"
	"github.com/ximply/exporter-manger/config"
	"strings"
        "fmt"
)

func overFilters(src string, filters map[string]string) string {
	ret := ""
	metrics := strings.Split(src, "\n")
	for _, s := range metrics {
		if strings.HasPrefix(s, "#") {
			continue
		}

		tmp := utils.Substr(s, 0, strings.Index(s, " "))
		metric := tmp
		if strings.Contains(tmp, "{") {
			metric = utils.Substr(tmp, 0, strings.Index(tmp, "{"))
		}
		fmt.Println(metric)
		if utils.ExistsMetric(metric, filters) {
			ret += s
			ret += "\n"
		}
	}

	return ret
}

func NodeMetrics() string {
	rsp := utils.MetricsFromUnixSock(config.NodeConfig().BaseCfg.UnixSockFile,
		config.NodeConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.NodeConfig().BaseCfg.Filters)
}

func RedisMetrics() string {
	rsp := utils.MetricsFromUnixSock(config.RedisConfig().BaseCfg.UnixSockFile,
		config.RedisConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.RedisConfig().BaseCfg.Filters)
}
