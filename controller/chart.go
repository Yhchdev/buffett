package controller

import (
	"fmt"
	"github.com/Yhchdev/buffett/datacenter/eastmoney"
	"github.com/Yhchdev/buffett/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strings"
)

const maxHistory = 9

type Char struct {
	Name   string   `json:"name"`
	Series []Series `json:"series"`
}

type Series struct {
	Type string        `json:"type"`
	Name string        `json:"name"`
	X    []interface{} `json:"x"`
	Y    []interface{} `json:"y"`
}

func Chart(c *gin.Context) {

	upperStr := strings.ToUpper(c.Query("secucode"))
	reportType := c.Query("report")
	// 获取数据
	history, err := eastmoney.NewEastMoney().QueryHistoricalFinaMainData(c, upperStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	usefulHistory := history.FilterByReportType(eastmoney.FinaReportType(reportType))

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	tableX := []interface{}{}
	KCFJCXSYJLR := []interface{}{}
	KCFJCXSYJLRTZ := []interface{}{}
	Totaloperatereve := []interface{}{}
	Totaloperaterevetz := []interface{}{}

	if len(usefulHistory) > maxHistory {
		usefulHistory = usefulHistory[len(usefulHistory)-9:]
	}

	for _, item := range usefulHistory {

		tableX = append(tableX, item.ReportYear)
		KCFJCXSYJLR = append(KCFJCXSYJLR, utils.ConvertToBillions(item.Kcfjcxsyjlr))
		KCFJCXSYJLRTZ = append(KCFJCXSYJLRTZ, utils.FloatFormat(item.Kcfjcxsyjlrtz))
		Totaloperatereve = append(Totaloperatereve, utils.ConvertToBillions(item.Totaloperatereve))
		Totaloperaterevetz = append(Totaloperaterevetz, utils.FloatFormat(item.Totaloperaterevetz))
	}

	charts := make([]Char, 0)

	charts = append(charts, Char{
		Name: "扣非净利润及其增长率",
		Series: []Series{
			{
				Name: "扣非净利润",
				Type: "bar",
				X:    tableX,
				Y:    KCFJCXSYJLR,
			},
			{
				Name: "增长率",
				Type: "line",
				X:    tableX,
				Y:    KCFJCXSYJLRTZ,
			},
		},
	}, Char{
		Name: "营业收入及其增长率",
		Series: []Series{
			{
				Name: "营业收入",
				Type: "bar",
				X:    tableX,
				Y:    Totaloperatereve,
			},
			{
				Name: "增长率",
				Type: "line",
				X:    tableX,
				Y:    Totaloperaterevetz,
			},
		},
	})

	// 处理请求并返回响应
	c.JSON(200, gin.H{
		"charts": charts,
	})

}
