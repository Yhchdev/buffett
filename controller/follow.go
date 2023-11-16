package controller

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Message struct {
	FromUserName string `xml:"FromUserName"`
}

type FollowReq struct {
	XmlData string
}

type FollowResp struct {
}

func Follow(c *gin.Context) {

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

	// todo save to db

	fmt.Println("FromUserName:", message.FromUserName)

}
