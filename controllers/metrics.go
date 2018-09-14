package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ximply/exporter-manger/models"
	"github.com/ximply/exporter-manger/config"
	"strings"
	"github.com/astaxie/beego/context"
	"fmt"
	"regexp"
)

// response to prometheus
func response(r string, ctx *context.Context) {
	if strings.Contains(r, "_") || strings.Contains(r, "{") {
		ctx.Output.Body([]byte(r))
		return
	}
	ctx.Output.SetStatus(502)
	ctx.Output.Body([]byte(``))
}

// company
type CompanyController struct {
	beego.Controller
}

func (c *CompanyController) CompanyMetrics() {
	r := models.CompanyMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// company info
type CompanyInfoController struct {
	beego.Controller
}

func (c *CompanyInfoController) CompanyInfoMetrics() {
	r := models.CompanyInfoMetrics()
	c.Ctx.Output.Body([]byte(r))
}

type CompanyInfoJsonController struct {
	beego.Controller
}

func (c *CompanyInfoJsonController) CompanyInfoJsonMetrics() {
	r := models.CompanyInfoJsonMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// company heartbeat
type CompanyHbController struct {
	beego.Controller
}

func (c *CompanyHbController) CompanyHbMetrics() {
	r := models.CompanyHbMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// company conn
type CompanyConnController struct {
	beego.Controller
}

func (c *CompanyConnController) CompanyConnMetrics() {
	r := models.CompanyConnMetrics()
	c.Ctx.Output.Body([]byte(r))
}

// heartbeat
type HeartbeatController struct {
	beego.Controller
}

func (c *HeartbeatController) Heartbeat() {
	c.Ctx.Output.Body([]byte("heartbeat 1\n"))
}

// node exporter
type NodeController struct {
	beego.Controller
}

func (c *NodeController) NodeMetrics() {
	response(models.NodeMetrics(), c.Ctx)
}

// nginx exporter
type NginxController struct {
	beego.Controller
}

func (c *NginxController) NginxMetrics() {
	response(models.NginxMetrics(), c.Ctx)
}

// nginx vts exporter
type NginxVtsController struct {
	beego.Controller
}

func (c *NginxVtsController) NginxVtsMetrics() {
	response(models.NginxVtsMetrics(), c.Ctx)
}

// php-fpm exporter
type PhpfpmController struct {
	beego.Controller
}

func (c *PhpfpmController) PhpfpmMetrics() {
	response(models.PhpfpmMetrics(), c.Ctx)
}

// redis exporter
type RedisController struct {
	beego.Controller
}

func (c *RedisController) RedisMetrics() {
	response(models.RedisMetrics(), c.Ctx)
}

// memcached exporter
type MemchachedController struct {
	beego.Controller
}

func (c *MemchachedController) MemcachedMetrics() {
	response(models.MemcachedMetrics(), c.Ctx)
}

// mysqld exporter
type MysqldController struct {
	beego.Controller
}

func (c *MysqldController) MysqldMetrics() {
	response(models.MysqldMetrics(), c.Ctx)
}

// multi mysqld exporter
type MultiMysqldController struct {
	beego.Controller
}

func (c *MultiMysqldController) MultiMysqldMetrics() {
	response(models.MultiMysqldMetrics(), c.Ctx)
}

// haproxy exporter
type HaproxyController struct {
	beego.Controller
}

func (c *HaproxyController) HaproxyMetrics() {
	response(models.HaproxyMetrics(), c.Ctx)
}

// gearman exporter
type GearmanController struct {
	beego.Controller
}

func (c *GearmanController) GearmanMetrics() {
	response(models.GearmanMetrics(), c.Ctx)
}

// mongodb exporter
type MongodbController struct {
	beego.Controller
}

func (c *MongodbController) MongodbMetrics() {
	response(models.MongodbMetrics(), c.Ctx)
}

// dellhardware exporter
type DellHardwareController struct {
	beego.Controller
}

func (c *DellHardwareController) DellHardwareMetrics() {
	response(models.DellHardwareMetrics(), c.Ctx)
}

// xenserver exporter
type XenserverController struct {
	beego.Controller
}

func (c *XenserverController) XenserverMetrics() {
	response(models.XenserverMetrics(), c.Ctx)
}

// elasticsearch exporter
type ElasticsearchController struct {
	beego.Controller
}

func (c *ElasticsearchController) ElasticsearchMetrics() {
	response(models.ElasticsearchMetrics(), c.Ctx)
}

// logstash exporter
type LogstashController struct {
	beego.Controller
}

func (c *LogstashController) LogstashMetrics() {
	response(models.LogstashMetrics(), c.Ctx)
}

// multi logstash exporter
type MultiLogstashController struct {
	beego.Controller
}

func (c *MultiLogstashController) MultiLogstashMetrics() {
	response(models.MultiLogstashMetrics(), c.Ctx)
}

// ping exporter
type PingController struct {
	beego.Controller
}

func (c *PingController) PingMetrics() {
	response(models.PingMetrics(), c.Ctx)
}

// tcp ping exporter
type TcpPingController struct {
	beego.Controller
}

func (c *TcpPingController) TcpPingMetrics() {
	response(models.TcpPingMetrics(), c.Ctx)
}

// httpstat exporter
type HttpStatController struct {
	beego.Controller
}

func (c *HttpStatController) HttpStatMetrics() {
	response(models.HttpStatMetrics(), c.Ctx)
}

// ping domain exporter
type PingDomainController struct {
	beego.Controller
}

func (c *PingDomainController) PingDomainMetrics() {
	response(models.PingDomainMetrics(), c.Ctx)
}

// certwacher exporter
type CertwacherController struct {
	beego.Controller
}

func (c *CertwacherController) CertwacherMetrics() {
	response(models.CertwacherMetrics(), c.Ctx)
}

// alive exporter
type AliveController struct {
	beego.Controller
}

func (c *AliveController) AliveMetrics() {
	r := models.AliveMetrics()
	chkSize := len(config.AliveConfig().Checks)
	if strings.Contains(r, "_") || strings.Contains(r, "{") {
		if chkSize == 0 {
			c.Ctx.Output.Body([]byte("alive_up 1\n"))
			return
		}

		l := strings.Split(r, "\n")
		ret := ""
		p := regexp.MustCompile(`\"([^\"]*)\"`)
		//fmt.Println(p.FindAllString(`alive{type="process",pname="expmgr",pargs="null"} 1`, -1))
		checks := config.AliveConfig().Checks
		for _, c := range checks {
			pRetArgs := ""
			alive := false
			for _, i := range l {
				if !strings.HasPrefix(i, `alive{type="process"`) {
					continue
				}
				pl := p.FindAllString(i, -1)
				if len(pl) != 3 {
					continue
				}
				pName := pl[1]
				pName = strings.Replace(pName, "\"", "", -1)
				pArgs := pl[2]
				// check process name
				if strings.Compare(pName, c.ProcessName) != 0 {
					continue
				}
				alive = true
				// check args
				for _, arg := range c.KeyList {
					if !strings.Contains(pArgs, arg) {
						alive = false
						break
					}
				}
				keyListSize := len(c.KeyList)
				if keyListSize == 0 {
					alive = true
				}

				if !alive {
					continue
				}

				pRetArgs = pArgs
			}

			if alive {
				ret += fmt.Sprintf("alive{type=\"process\",pname=\"%s\",pargs=%s,alias=\"%s\"} 1\n",
					c.ProcessName, pRetArgs, c.Alias)
			} else {
				ret += fmt.Sprintf("alive{type=\"process\",pname=\"%s\",pargs=\"%s\",alias=\"%s\"} 0\n",
					c.ProcessName, "null", c.Alias)
			}
		}
		ret += "alive_up 1\n"
		c.Ctx.Output.Body([]byte(ret))
		return
	}

	c.Ctx.Output.Body([]byte("alive_up 0\n"))
}

// rabbitmq exporter
type RabbitmqController struct {
	beego.Controller
}

func (c *RabbitmqController) RabbitmqMetrics() {
	response(models.RabbitmqMetrics(), c.Ctx)
}

// supervisor exporter
type SupervisorController struct {
	beego.Controller
}

func (c *SupervisorController) SupervisorMetrics() {
	response(models.SupervisorMetrics(), c.Ctx)
}

// java exporter
type JavaController struct {
	beego.Controller
}

func (c *JavaController) JavaMetrics() {
	response(models.JavaMetrics(), c.Ctx)
}

type JavaInfoController struct {
	beego.Controller
}

func (c *JavaInfoController) JavaInfo() {
	//body, _ := ioutil.ReadAll(c.Ctx.Request.Body)
	bodyStr := string(c.Ctx.Input.RequestBody)
    models.JavaInfo(bodyStr)

	out := make(map[string]interface{})
	out["status"] = 0
	out["msg"] = ""
	c.Data["json"] = out
	c.ServeJSON()
}

// beanstalkd exporter
type BeanstalkdController struct {
	beego.Controller
}

func (c *BeanstalkdController) BeanstalkdMetrics() {
	response(models.BeanstalkdMetrics(), c.Ctx)
}

// bind exporter
type BindController struct {
	beego.Controller
}

func (c *BindController) BindMetrics() {
	response(models.BindMetrics(), c.Ctx)
}

// solr exporter
type SolrController struct {
	beego.Controller
}

func (c *SolrController) SolrMetrics() {
	response(models.SolrMetrics(), c.Ctx)
}

// hadoop data node exporter
type HadoopDataNodeController struct {
	beego.Controller
}

func (c *HadoopDataNodeController) HadoopDataNodeMetrics() {
	response(models.HadoopDataNodeMetrics(), c.Ctx)
}

// hadoop name node exporter
type HadoopNameNodeController struct {
	beego.Controller
}

func (c *HadoopNameNodeController) HadoopNameNodeMetrics() {
	response(models.HadoopNameNodeMetrics(), c.Ctx)
}

// hadoop second name node exporter
type HadoopSecondNameNodeController struct {
	beego.Controller
}

func (c *HadoopSecondNameNodeController) HadoopSecondNameNodeMetrics() {
	response(models.HadoopSecondNameNodeMetrics(), c.Ctx)
}

// hadoop resource manager exporter
type HadoopResourceManagerController struct {
	beego.Controller
}

func (c *HadoopResourceManagerController) HadoopResourceManagerMetrics() {
	response(models.HadoopResourceManagerMetrics(), c.Ctx)
}

// kafka exporter
type KafkaController struct {
	beego.Controller
}

func (c *KafkaController) KafkaMetrics() {
	response(models.KafkaMetrics(), c.Ctx)
}

// zookeeper exporter
type ZookeeperController struct {
	beego.Controller
}

func (c *ZookeeperController) ZookeeperMetrics() {
	response(models.ZookeeperMetrics(), c.Ctx)
}