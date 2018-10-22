package uploads

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"github.com/ximply/exporter-manger/models"
)

func JavaInfo(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, `{"ret":-1,"msg":"read request body error"}`)
		return
	}
	models.JavaInfo(string(body))

	out := make(map[string]interface{})
	out["status"] = 0
	out["msg"] = ""
	c.SecureJSON(200, out)
}
