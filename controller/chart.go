package controller

import (
	"fmt"
	"github.com/Yhchdev/buffett/datacenter/eastmoney"
	"github.com/Yhchdev/buffett/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"strings"
)

type Cha struct {
	Type string        `json:"type"`
	X    []interface{} `json:"x"`
	Y    []interface{} `json:"y"`
}

func Chart(c *gin.Context) {

	upperStr := strings.ToUpper(c.Query("secucode"))

	fmt.Println(upperStr)

	// 获取数据
	history, err := eastmoney.NewEastMoney().QueryHistoricalFinaMainData(c, upperStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	usefulHistory := history.FilterByReportType(eastmoney.FinaReportTypeQ3)

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

	for _, item := range usefulHistory {

		tableX = append(tableX, item.ReportYear)
		KCFJCXSYJLR = append(KCFJCXSYJLR, utils.ConvertToBillions(item.Kcfjcxsyjlr))
		KCFJCXSYJLRTZ = append(KCFJCXSYJLRTZ, utils.FloatFormat(item.Kcfjcxsyjlrtz))
		Totaloperatereve = append(Totaloperatereve, utils.ConvertToBillions(item.Totaloperatereve))
		Totaloperaterevetz = append(Totaloperaterevetz, utils.FloatFormat(item.Totaloperaterevetz))
	}

	charts := make([][]Cha, 0)

	charts = append(charts, []Cha{
		{
			Type: "bar",
			X:    tableX,
			Y:    KCFJCXSYJLR,
		},
		{
			Type: "line",
			X:    tableX,
			Y:    KCFJCXSYJLRTZ,
		},
	})

	// 处理请求并返回响应
	c.JSON(200, gin.H{
		"charts": charts,
	})

}
