package model

import (
	"time"
)

type User struct {
	Id           int64     `json:"id"`
	UserName     string    `json:"user_name"` // shiniugu_ + open_id[:6]
	OpenId       string    `json:"open_id"`
	VipType      string    `json:"vip_type"`      // 0: 普通用户 1: 日会员用户 2.月会员 3.年会员 4.终身会员
	TotalCost    int64     `json:"total_cost"`    // 总消费金额 累加充值金额 用于计算会员等级
	LastRecharge time.Time `json:"last_recharge"` // 最后充值时间
	IsFollow     bool      `json:"is_follow"`     // 是否还在关注着公众号
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
