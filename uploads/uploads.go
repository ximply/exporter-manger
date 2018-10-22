package uploads

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ximply/exporter-manger/config"
	"github.com/yireyun/go-queue"
)

var gUploadExceptionQueue *queue.EsQueue

func Index(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func ValidGroup(group string) bool {
	for _, i := range config.UConfigs().Groups {
		if group == i {
			return true
		}
	}

	return false
}

func ExceptionQueue() *queue.EsQueue {
	return gUploadExceptionQueue
}

func Start() {
	go func() {
		Init().Run("127.0.0.1:" + config.UConfigs().ListenPort)
	}()
}

func Init() *gin.Engine {
	gUploadExceptionQueue = queue.NewQueue(uint32(config.UConfigs().UploadQueueSize))
	StartUploadException()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/", Index)
	router.POST("/exception", Exception)
	router.POST("/javainfo", JavaInfo)

	return router
}
