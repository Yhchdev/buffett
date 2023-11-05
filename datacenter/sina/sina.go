// Package sina 新浪财经接口封装
package sina

import (
	"github.com/go-resty/resty/v2"
)

// Sina 新浪财经数据源
type Sina struct {
	// http 客户端
	HTTPClient *resty.Client
}

// NewSina 创建 Sina 实例
func NewSina() Sina {

	return Sina{
		HTTPClient: resty.New(),
	}
}
