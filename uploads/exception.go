package uploads

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/json-iterator/go"
	"github.com/ximply/exporter-manger/structs"
	"github.com/ximply/exporter-manger/config"
	"time"
	"github.com/parnurzeal/gorequest"
)

func Exception(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, `{"ret":-1,"msg":"read request body error"}`)
		return
	}

	ex := structs.ExceptionItem{}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	if json.Unmarshal(body, &ex) != nil {
		c.JSON(http.StatusOK, `{"ret":-1,"msg":"parse body fail[json]"}`)
		return
	}

	if !ValidGroup(ex.Group) {
		c.JSON(http.StatusOK, `{"ret":-1,"msg":"error group"}`)
		return
	}

	ExceptionQueue().Put(&ex)
	c.JSON(http.StatusOK, `{"ret":0,"msg":""}`)
}

func packExceptionData() {
	its := make([]interface{}, config.UConfigs().UploadQueueSize)
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	for {
		gets, _ := ExceptionQueue().Gets(its)
		count := int(gets)
		if count > 0 {
			d, err := json.Marshal(its[0 : count - 1])
			if err == nil {
				go exceptionToAlertCenter(d)
			}
		}
		time.Sleep(time.Duration(config.UConfigs().UploadDelaySec) * time.Second)
	}
}

func exceptionToAlertCenter(d []byte) {
	req := gorequest.New()
	req.BounceToRawString = true
	req.Retry(config.UConfigs().Retry, time.Duration(config.UConfigs().TimeoutSec) * time.Second,
		http.StatusBadRequest, http.StatusInternalServerError).
		Post(config.UConfigs().AcUrl).SendString(string(d)).End()
}

func StartUploadException() {
	go packExceptionData()
}