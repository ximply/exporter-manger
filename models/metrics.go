package models

import (
	"github.com/ximply/exporter-manger/utils"
	"github.com/ximply/exporter-manger/config"
	"strings"
	"fmt"
	"net"
	"net/http"
	"io/ioutil"
	"context"
)

type UnixResponse struct {
	Rsp string
	Status string
}

func existsMetric(metric string, filters map[string]string) bool {
	if filters == nil {
		return true
	}

	if len(filters) == 0 {
		return true
	}

	if _, ok := filters[metric]; ok {
		return true
	}
	return false
}

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
		if existsMetric(metric, filters) {
			ret += s
			ret += "\n"
		}
	}

	return ret
}

func metricsFromUnixSock(unixSockFile string, metricsPath string) UnixResponse {
	rsp := UnixResponse{
		Rsp: "",
		Status: "500",
	}
	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", unixSockFile)
			},
		},
		Timeout: config.NodeConfig().BaseCfg.Timeout,
	}
	res, err := c.Get(fmt.Sprintf("http://unix/%s", metricsPath))
	if err != nil {
		res.Body.Close()
		return rsp
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return rsp
	}
	rsp.Rsp = string(body)
	rsp.Status = "200"
	return rsp
}

func NodeMetrics() string {
	rsp := metricsFromUnixSock(config.NodeConfig().BaseCfg.UnixSockFile,
		config.NodeConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.NodeConfig().BaseCfg.Filters)
}

func NginxMetrics() string {
	rsp := metricsFromUnixSock(config.NginxConfig().BaseCfg.UnixSockFile,
		config.NginxConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.NginxConfig().BaseCfg.Filters)
}

func NginxVtsMetrics() string {
	rsp := metricsFromUnixSock(config.NginxVtsConfig().BaseCfg.UnixSockFile,
		config.NginxVtsConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.NginxVtsConfig().BaseCfg.Filters)
}

func PhpfpmMetrics() string {
	rsp := metricsFromUnixSock(config.PhpfpmConfig().BaseCfg.UnixSockFile,
		config.PhpfpmConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.PhpfpmConfig().BaseCfg.Filters)
}

func RedisMetrics() string {
	rsp := metricsFromUnixSock(config.RedisConfig().BaseCfg.UnixSockFile,
		config.RedisConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.RedisConfig().BaseCfg.Filters)
}

func MemcachedMetrics() string {
	rsp := metricsFromUnixSock(config.MemcachedConfig().BaseCfg.UnixSockFile,
		config.MemcachedConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.MemcachedConfig().BaseCfg.Filters)
}

func MysqldMetrics() string {
	rsp := metricsFromUnixSock(config.MysqlConfig().BaseCfg.UnixSockFile,
		config.MysqlConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.MysqlConfig().BaseCfg.Filters)
}

func HaproxyMetrics() string {
	rsp := metricsFromUnixSock(config.HaproxyConfig().BaseCfg.UnixSockFile,
		config.HaproxyConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.HaproxyConfig().BaseCfg.Filters)
}

func GearmanMetrics() string {
	rsp := metricsFromUnixSock(config.GearmanConfig().BaseCfg.UnixSockFile,
		config.GearmanConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.GearmanConfig().BaseCfg.Filters)
}

func MongodbMetrics() string {
	rsp := metricsFromUnixSock(config.MongodbConfig().BaseCfg.UnixSockFile,
		config.MongodbConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.MongodbConfig().BaseCfg.Filters)
}

func DellHardwareMetrics() string {
	rsp := metricsFromUnixSock(config.DellHardwareConfig().BaseCfg.UnixSockFile,
		config.DellHardwareConfig().BaseCfg.MetricsPath)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.DellHardwareConfig().BaseCfg.Filters)
}