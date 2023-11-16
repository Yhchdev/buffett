package controller

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UnFollowReq struct {
	XmlData string
}

type UnFollowResp struct {
}

func Unfollow(c *gin.Context) {

	followReq := FollowReq{}

	err := c.ShouldBind(followReq)
	if err != nil {
		return
	}

	var message Message
	if err := xml.Unmarshal([]byte(followReq.XmlData), &message); err != nil {
		fmt.Println("解析 XML 失败:", err)
		return
	}

	// 更新follow状态

	fmt.Println("FromUserName:", message.FromUserName)

}
