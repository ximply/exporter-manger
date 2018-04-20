package utils

import (
        "fmt"
	"io/ioutil"
	"context"
	"net"
	"net/http"
	"github.com/ximply/exporter-manger/config"
)

type UnixResponse struct {
	Rsp string
	Status string
}

func ExistsMetric(metric string, filters map[string]string) bool {
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

func MetricsFromUnixSock(unixSockFile string, metricsPath string) UnixResponse {
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

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}
