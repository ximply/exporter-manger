package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/models"
)

type NodeController struct {
	beego.Controller
}

func (c *NodeController) NodeMetrics() {
	r := models.NodeMetrics()
	c.Ctx.Output.Body([]byte(r))
}
