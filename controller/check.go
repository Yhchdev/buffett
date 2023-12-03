package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WxCheck(c *gin.Context) {

	// 首次校验
	echostr := c.Query("echostr")
	if echostr != "" {
		c.String(http.StatusOK, echostr)
	}

}
