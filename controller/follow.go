package controller

import (
	"encoding/xml"
	"fmt"
	"github.com/Yhchdev/buffett/config"
	"github.com/Yhchdev/buffett/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"io/ioutil"
	"log"
	"time"
)

// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Receiving_event_pushes.html

type Message struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"` // openid
	CreateTime   string `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Event        string `json:"Event"` // 事件类型，subscribe(订阅)、unsubscribe(取消订阅)
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

	db := config.DB

	now := time.Now()

	user := model.User{
		Id:        now.UnixNano(),
		UserName:  fmt.Sprintf("ShiNiuGu_%s", message.FromUserName[len(message.FromUserName)-6:]),
		OpenId:    message.FromUserName,
		TotalCost: 0,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if message.Event == "subscribe" {
		user.IsFollow = true
	} else {
		user.IsFollow = false
	}

	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "open_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"role", "is_follow"}),
	}).Create(&user)

	// todo 重定向到搜索页

}
