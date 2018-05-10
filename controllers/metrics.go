package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/models"
)


// node exporter
type NodeController struct {
	beego.Controller
}

func (c *NodeController) NodeMetrics() {
	r := models.NodeMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// nginx exporter
type NginxController struct {
	beego.Controller
}

func (c *NginxController) NginxMetrics() {
	r := models.NginxMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// nginx vts exporter
type NginxVtsController struct {
	beego.Controller
}

func (c *NginxVtsController) NginxVtsMetrics() {
	r := models.NginxVtsMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// php-fpm exporter
type PhpfpmController struct {
	beego.Controller
}

func (c *PhpfpmController) PhpfpmMetrics() {
	r := models.PhpfpmMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// redis exporter
type RedisController struct {
	beego.Controller
}

func (c *RedisController) RedisMetrics() {
	r := models.RedisMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// memcached exporter
type MemchachedController struct {
	beego.Controller
}

func (c *MemchachedController) MemcachedMetrics() {
	r := models.MemcachedMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// mysqld exporter
type MysqldController struct {
	beego.Controller
}

func (c *MysqldController) MysqldMetrics() {
	r := models.MysqldMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// multi mysqld exporter
type MultiMysqldController struct {
	beego.Controller
}

func (c *MultiMysqldController) MultiMysqldMetrics() {
	r := models.MultiMysqldMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// haproxy exporter
type HaproxyController struct {
	beego.Controller
}

func (c *HaproxyController) HaproxyMetrics() {
	r := models.HaproxyMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// gearman exporter
type GearmanController struct {
	beego.Controller
}

func (c *GearmanController) GearmanMetrics() {
	r := models.GearmanMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// mongodb exporter
type MongodbController struct {
	beego.Controller
}

func (c *MongodbController) MongodbMetrics() {
	r := models.MongodbMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// dellhardware exporter
type DellHardwareController struct {
	beego.Controller
}

func (c *DellHardwareController) DellHardwareMetrics() {
	r := models.DellHardwareMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// xenserver exporter
type XenserverController struct {
	beego.Controller
}

func (c *XenserverController) XenserverMetrics() {
	r := models.XenserverMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// elasticsearch exporter
type ElasticsearchController struct {
	beego.Controller
}

func (c *ElasticsearchController) ElasticsearchMetrics() {
	r := models.ElasticsearchMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// logstash exporter
type LogstashController struct {
	beego.Controller
}

func (c *LogstashController) LogstashMetrics() {
	r := models.LogstashMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// ping exporter
type PingController struct {
	beego.Controller
}

func (c *PingController) PingMetrics() {
	r := models.PingMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// tcp ping exporter
type TcpPingController struct {
	beego.Controller
}

func (c *TcpPingController) TcpPingMetrics() {
	r := models.TcpPingMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// httpstat exporter
type HttpStatController struct {
	beego.Controller
}

func (c *HttpStatController) HttpStatMetrics() {
	r := models.HttpStatMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// ping domain exporter
type PingDomainController struct {
	beego.Controller
}

func (c *PingDomainController) PingDomainMetrics() {
	r := models.PingDomainMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// certwacher exporter
type CertwacherController struct {
	beego.Controller
}

func (c *CertwacherController) CertwacherMetrics() {
	r := models.CertwacherMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// alive exporter
type AliveController struct {
	beego.Controller
}

func (c *AliveController) AliveMetrics() {
	r := models.AliveMetrics()
	c.Ctx.Output.Body([]byte(r))
}