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

// redis exporter
type RedisController struct {
	beego.Controller
}

func (c *NodeController) RedisMetrics() {
	r := models.RedisMetrics()
	c.Ctx.Output.Body([]byte(r))
}