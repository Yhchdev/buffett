package controller

import (
	"github.com/Yhchdev/buffett/datacenter/sina"
	"github.com/gin-gonic/gin"
)

const maxNum = 6

type Stock struct {
	// 带后缀的代码
	Secucode string `json:"value"`
	// 股票名称
	Name string `json:"label"`
}

func Search(c *gin.Context) {
	keyword := c.Query("keyword")

	stocks, err := sina.NewSina().KeywordSearch(c, keyword)
	if err != nil {
		return
	}

	stocksResp := make([]Stock, 0)

	for i, stock := range stocks {
		stocksResp = append(stocksResp, Stock{
			Secucode: stock.Secucode,
			Name:     stock.Name,
		})
		if i+1 >= maxNum {
			break
		}
	}

	c.JSON(200, gin.H{
		"stocks": stocksResp,
	})

}
