// 东方财富数据源封装

package eastmoney

import (
	"github.com/go-resty/resty/v2"
)

// EastMoney 东方财富数据源
type EastMoney struct {
	// http 客户端
	HTTPClient *resty.Client
}

// NewEastMoney 创建 EastMoney 实例
func NewEastMoney() EastMoney {
	return EastMoney{
		HTTPClient: resty.New(),
	}
}
