package controller

import (
	"github.com/Yhchdev/buffett/datacenter/eastmoney"
	"github.com/gin-gonic/gin"
)

func HotStock(c *gin.Context) {

	stocks, err := eastmoney.NewEastMoney().Xuangu(c)
	if err != nil {
		return
	}

	stocksResp := make([]Stock, 0)

	for i, stock := range stocks {
		stocksResp = append(stocksResp, Stock{
			Secucode: stock.SECUCODE,
			Name:     stock.SECURITYNAMEABBR,
		})
		if i+1 >= 50 {
			break
		}
	}

	c.JSON(200, gin.H{
		"stocks": stocksResp,
	})

}
