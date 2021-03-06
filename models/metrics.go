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
	"time"
	"sync"
	"runtime"
)

type UnixResponse struct {
	Rsp string
	Status string
}

type JavaInfoItem struct {
	Info string
	UpdateTime time.Time
}
var g_javaInfoLock sync.RWMutex
var g_javInfo map[string]*JavaInfoItem
var g_numCpuCores float64

func Init() {
	g_javInfo = make(map[string]*JavaInfoItem)
    g_numCpuCores = float64(runtime.NumCPU())
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
		if existsMetric(metric, filters) {
			ret += s
			ret += "\n"
		}
	}

	return ret
}

func metricsFromUnixSock(unixSockFile string, metricsPath string, timeout time.Duration) UnixResponse {
	rsp := UnixResponse{
		Rsp: "",
		Status: "500",
	}
	c := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", unixSockFile)
			},
			DisableKeepAlives: true,
		},
		Timeout: timeout,
	}

	res, err := c.Get(fmt.Sprintf("http://unix/%s", metricsPath))
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return rsp
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return rsp
	}
	rsp.Rsp = string(body)
	rsp.Status = "200"
	return rsp
}

func CompanyMetrics() string {
	rsp := metricsFromUnixSock(config.CompanyConfig().BaseCfg.UnixSockFile,
		config.CompanyConfig().BaseCfg.MetricsPath,
		config.CompanyConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.CompanyConfig().BaseCfg.Filters)
}

func CompanyInfoMetrics() string {
	rsp := metricsFromUnixSock(config.CompanyInfoConfig().BaseCfg.UnixSockFile,
		config.CompanyInfoConfig().BaseCfg.MetricsPath,
		config.CompanyInfoConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.CompanyInfoConfig().BaseCfg.Filters)
}

func CompanyInfoJsonMetrics() string {
	rsp := metricsFromUnixSock(config.CompanyConfig().BaseCfg.UnixSockFile,
		"/json",
		config.CompanyConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return rsp.Rsp
}

func CompanyHbMetrics() string {
	rsp := metricsFromUnixSock(config.CompanyHbConfig().BaseCfg.UnixSockFile,
		config.CompanyHbConfig().BaseCfg.MetricsPath,
		config.CompanyHbConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.CompanyHbConfig().BaseCfg.Filters)
}

func CompanyConnMetrics() string {
	rsp := metricsFromUnixSock(config.CompanyConnConfig().BaseCfg.UnixSockFile,
		config.CompanyConnConfig().BaseCfg.MetricsPath,
		config.CompanyConnConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.CompanyConnConfig().BaseCfg.Filters)
}

func NodeMetrics() string {
	rsp := metricsFromUnixSock(config.NodeConfig().BaseCfg.UnixSockFile,
		config.NodeConfig().BaseCfg.MetricsPath,
		config.NodeConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}
	rsp.Rsp = fmt.Sprintf("node_num_cpu_cores %g\n", g_numCpuCores) + rsp.Rsp
	return overFilters(rsp.Rsp, config.NodeConfig().BaseCfg.Filters)
}

func NginxMetrics() string {
	rsp := metricsFromUnixSock(config.NginxConfig().BaseCfg.UnixSockFile,
		config.NginxConfig().BaseCfg.MetricsPath,
		config.NginxConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.NginxConfig().BaseCfg.Filters)
}

func NginxVtsMetrics() string {
	rsp := metricsFromUnixSock(config.NginxVtsConfig().BaseCfg.UnixSockFile,
		config.NginxVtsConfig().BaseCfg.MetricsPath,
		config.NginxVtsConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.NginxVtsConfig().BaseCfg.Filters)
}

func PhpfpmMetrics() string {
	rsp := metricsFromUnixSock(config.PhpfpmConfig().BaseCfg.UnixSockFile,
		config.PhpfpmConfig().BaseCfg.MetricsPath,
		config.PhpfpmConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.PhpfpmConfig().BaseCfg.Filters)
}

func RedisMetrics() string {
	rsp := metricsFromUnixSock(config.RedisConfig().BaseCfg.UnixSockFile,
		config.RedisConfig().BaseCfg.MetricsPath,
		config.RedisConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.RedisConfig().BaseCfg.Filters)
}

func Redis2Metrics() string {
	rsp := metricsFromUnixSock(config.Redis2Config().BaseCfg.UnixSockFile,
		config.Redis2Config().BaseCfg.MetricsPath,
		config.Redis2Config().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.Redis2Config().BaseCfg.Filters)
}

func MemcachedMetrics() string {
	rsp := metricsFromUnixSock(config.MemcachedConfig().BaseCfg.UnixSockFile,
		config.MemcachedConfig().BaseCfg.MetricsPath,
		config.MemcachedConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.MemcachedConfig().BaseCfg.Filters)
}

func MysqldMetrics() string {
	rsp := metricsFromUnixSock(config.MysqlConfig().BaseCfg.UnixSockFile,
		config.MysqlConfig().BaseCfg.MetricsPath,
		config.MysqlConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.MysqlConfig().BaseCfg.Filters)
}

func MultiMysqldMetrics() string {
	rsp := metricsFromUnixSock(config.MultiMysqlConfig().BaseCfg.UnixSockFile,
		config.MultiMysqlConfig().BaseCfg.MetricsPath,
		config.MultiMysqlConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.MultiMysqlConfig().BaseCfg.Filters)
}

func HaproxyMetrics() string {
	rsp := metricsFromUnixSock(config.HaproxyConfig().BaseCfg.UnixSockFile,
		config.HaproxyConfig().BaseCfg.MetricsPath,
		config.HaproxyConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.HaproxyConfig().BaseCfg.Filters)
}

func GearmanMetrics() string {
	rsp := metricsFromUnixSock(config.GearmanConfig().BaseCfg.UnixSockFile,
		config.GearmanConfig().BaseCfg.MetricsPath,
		config.GearmanConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.GearmanConfig().BaseCfg.Filters)
}

func MongodbMetrics() string {
	rsp := metricsFromUnixSock(config.MongodbConfig().BaseCfg.UnixSockFile,
		config.MongodbConfig().BaseCfg.MetricsPath,
		config.MongodbConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.MongodbConfig().BaseCfg.Filters)
}

func DellHardwareMetrics() string {
	rsp := metricsFromUnixSock(config.DellHardwareConfig().BaseCfg.UnixSockFile,
		config.DellHardwareConfig().BaseCfg.MetricsPath,
		config.DellHardwareConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.DellHardwareConfig().BaseCfg.Filters)
}

func XenserverMetrics() string {
	rsp := metricsFromUnixSock(config.XenserverConfig().BaseCfg.UnixSockFile,
		config.XenserverConfig().BaseCfg.MetricsPath,
		config.XenserverConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.XenserverConfig().BaseCfg.Filters)
}

func ElasticsearchMetrics() string {
	rsp := metricsFromUnixSock(config.ElasticsearchConfig().BaseCfg.UnixSockFile,
		config.ElasticsearchConfig().BaseCfg.MetricsPath,
		config.ElasticsearchConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.ElasticsearchConfig().BaseCfg.Filters)
}

func LogstashMetrics() string {
	rsp := metricsFromUnixSock(config.LogstashConfig().BaseCfg.UnixSockFile,
		config.LogstashConfig().BaseCfg.MetricsPath,
		config.LogstashConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.LogstashConfig().BaseCfg.Filters)
}

func MultiLogstashMetrics() string {
	rsp := metricsFromUnixSock(config.MultiLogstashConfig().BaseCfg.UnixSockFile,
		config.MultiLogstashConfig().BaseCfg.MetricsPath,
		config.MultiLogstashConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.MultiLogstashConfig().BaseCfg.Filters)
}

func PingMetrics() string {
	rsp := metricsFromUnixSock(config.PingConfig().BaseCfg.UnixSockFile,
		config.PingConfig().BaseCfg.MetricsPath,
		config.PingConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.PingConfig().BaseCfg.Filters)
}

func TcpPingMetrics() string {
	rsp := metricsFromUnixSock(config.TcpPingConfig().BaseCfg.UnixSockFile,
		config.TcpPingConfig().BaseCfg.MetricsPath,
		config.TcpPingConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.TcpPingConfig().BaseCfg.Filters)
}

func HttpStatMetrics() string {
	rsp := metricsFromUnixSock(config.HttpStatConfig().BaseCfg.UnixSockFile,
		config.HttpStatConfig().BaseCfg.MetricsPath,
		config.HttpStatConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.HttpStatConfig().BaseCfg.Filters)
}

func PingDomainMetrics() string {
	rsp := metricsFromUnixSock(config.PingDomainConfig().BaseCfg.UnixSockFile,
		config.PingDomainConfig().BaseCfg.MetricsPath,
		config.PingDomainConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.PingDomainConfig().BaseCfg.Filters)
}

func CertwacherMetrics() string {
	rsp := metricsFromUnixSock(config.CertwacherConfig().BaseCfg.UnixSockFile,
		config.CertwacherConfig().BaseCfg.MetricsPath,
		config.CertwacherConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.CertwacherConfig().BaseCfg.Filters)
}

func AliveMetrics() string {
	rsp := metricsFromUnixSock(config.AliveConfig().BaseCfg.UnixSockFile,
		config.AliveConfig().BaseCfg.MetricsPath,
		config.AliveConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.AliveConfig().BaseCfg.Filters)
}

func RabbitmqMetrics() string {
	rsp := metricsFromUnixSock(config.RabbitmqConfig().BaseCfg.UnixSockFile,
		config.RabbitmqConfig().BaseCfg.MetricsPath,
		config.RabbitmqConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.RabbitmqConfig().BaseCfg.Filters)
}

func SupervisorMetrics() string {
	rsp := metricsFromUnixSock(config.SupervisorConfig().BaseCfg.UnixSockFile,
		config.SupervisorConfig().BaseCfg.MetricsPath,
		config.SupervisorConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.SupervisorConfig().BaseCfg.Filters)
}

func JavaMetrics() string {
	g_javaInfoLock.RLock()
	r := g_javInfo
	g_javaInfoLock.RUnlock()

	ret := ""
	now := time.Now().Unix()
	for _, v := range r {
		if now - v.UpdateTime.Unix() > 60 {
			continue
		}
		if strings.HasSuffix(v.Info,"\n") {
			ret += v.Info
		} else {
			ret += v.Info + "\n"
		}
	}
	ret = strings.Replace(ret, `java_heap_bytes{type="committed",`, `java_heap_committed_bytes{`, -1)
	ret = strings.Replace(ret, `java_heap_bytes{type="init",`, `java_heap_init_bytes{`, -1)
	ret = strings.Replace(ret, `java_heap_bytes{type="used",`, `java_heap_used_bytes{`, -1)
	ret = strings.Replace(ret, `java_heap_bytes{type="max",`, `java_heap_max_bytes{`, -1)
	return overFilters(ret, config.JavaConfig().BaseCfg.Filters)
}

func JavaInfo(info string) {
	s := strings.TrimRight(info, "\n")
	l := strings.Split(s, "\n")
	if len(l) > 1 {
		target := strings.Split(l[0],"|")
		if len(target) != 2 {
			return
		}

		exist := false
		g_javaInfoLock.RLock()
		if v, ok := g_javInfo[l[0]]; ok {
			v.Info = s
			v.UpdateTime = time.Now()
			exist = true
		}
		g_javaInfoLock.RUnlock()

		if !exist {
			i := &JavaInfoItem{
				Info:s,
				UpdateTime:time.Now(),
			}
			g_javaInfoLock.Lock()
			g_javInfo[l[0]] = i
			g_javaInfoLock.Unlock()
		}
	}
}

func BeanstalkdMetrics() string {
	rsp := metricsFromUnixSock(config.BeanstalkdConfig().BaseCfg.UnixSockFile,
		config.BeanstalkdConfig().BaseCfg.MetricsPath,
		config.BeanstalkdConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.BeanstalkdConfig().BaseCfg.Filters)
}

func BindMetrics() string {
	rsp := metricsFromUnixSock(config.BindConfig().BaseCfg.UnixSockFile,
		config.BindConfig().BaseCfg.MetricsPath,
		config.BindConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.BindConfig().BaseCfg.Filters)
}

func SolrMetrics() string {
	rsp := metricsFromUnixSock(config.SolrConfig().BaseCfg.UnixSockFile,
		config.SolrConfig().BaseCfg.MetricsPath,
		config.SolrConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}
	
	return overFilters(rsp.Rsp, config.SolrConfig().BaseCfg.Filters)
}

func HadoopDataNodeMetrics() string {
	rsp := metricsFromUnixSock(config.HadoopDataNodeConfig().BaseCfg.UnixSockFile,
		config.HadoopDataNodeConfig().BaseCfg.MetricsPath,
		config.HadoopDataNodeConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.HadoopDataNodeConfig().BaseCfg.Filters)
}

func HadoopNameNodeMetrics() string {
	rsp := metricsFromUnixSock(config.HadoopNameNodeConfig().BaseCfg.UnixSockFile,
		config.HadoopNameNodeConfig().BaseCfg.MetricsPath,
		config.HadoopNameNodeConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.HadoopNameNodeConfig().BaseCfg.Filters)
}

func HadoopSecondNameNodeMetrics() string {
	rsp := metricsFromUnixSock(config.HadoopSecondNameNodeConfig().BaseCfg.UnixSockFile,
		config.HadoopSecondNameNodeConfig().BaseCfg.MetricsPath,
		config.HadoopSecondNameNodeConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.HadoopSecondNameNodeConfig().BaseCfg.Filters)
}

func HadoopResourceManagerMetrics() string {
	rsp := metricsFromUnixSock(config.HadoopResourceManagerConfig().BaseCfg.UnixSockFile,
		config.HadoopResourceManagerConfig().BaseCfg.MetricsPath,
		config.HadoopResourceManagerConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.HadoopResourceManagerConfig().BaseCfg.Filters)
}

func KafkaMetrics() string {
	rsp := metricsFromUnixSock(config.KafkaConfig().BaseCfg.UnixSockFile,
		config.KafkaConfig().BaseCfg.MetricsPath,
		config.KafkaConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.KafkaConfig().BaseCfg.Filters)
}

func ZookeeperMetrics() string {
	rsp := metricsFromUnixSock(config.ZookeeperConfig().BaseCfg.UnixSockFile,
		config.ZookeeperConfig().BaseCfg.MetricsPath,
		config.ZookeeperConfig().BaseCfg.Timeout)
	if strings.Compare(rsp.Status, "200") != 0 {
		return ""
	}

	return overFilters(rsp.Rsp, config.ZookeeperConfig().BaseCfg.Filters)
}