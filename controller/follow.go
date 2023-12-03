package controller

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

type Message struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   string `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"`
	EventKey     string `json:"EventKey"`
	Ticket       string `json:"Ticket"`
}

type FollowReq struct {
	XmlData string
}

type FollowResp struct {
}

func Follow(c *gin.Context) {

	xmlData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		c.String(500, "Internal Server Error")
		return
	}

	var message Message
	if err := xml.Unmarshal(xmlData, &message); err != nil {
		fmt.Println("解析 XML 失败:", err)
		return
	}

	// todo save to db

	fmt.Println("FromUserName:", message.FromUserName)
	fmt.Println("FromUserName:", message.Event)

}
