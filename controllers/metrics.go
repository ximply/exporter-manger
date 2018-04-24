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