package main

import (
	"context"
	"fmt"
	"github.com/Yhchdev/buffett/controller"
	"github.com/Yhchdev/buffett/datacenter/eastmoney"
	"github.com/Yhchdev/buffett/job"
	"github.com/Yhchdev/buffett/middleware"
	"github.com/Yhchdev/buffett/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type Chart struct {
	Type string        `json:"type"`
	X    []interface{} `json:"x"`
	Y    []interface{} `json:"y"`
}

type FinaMainResp struct {
	Charts [][]Chart `json:"charts"`
}

func getData(c *gin.Context) {

	// 获取数据
	history, err := eastmoney.NewEastMoney().QueryHistoricalFinaMainData(context.Background(), "603259.SH")
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

	//

	//tableX := []interface{}{nil}
	//KCFJCXSYJLR := []interface{}{"扣非净利润(亿元)"}
	//KCFJCXSYJLRTZ := []interface{}{"扣非净利润同比增长(%)"}
	//Totaloperatereve := []interface{}{"营业收入(亿元)"}
	//Totaloperaterevetz := []interface{}{"营业收入同比增长(%)"}

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

	charts := make([][]Chart, 0)

	charts = append(charts, []Chart{
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

func main() {

	router := gin.Default()

	// 添加跨域中间件
	router.Use(middleware.Cors())

	// 定义路由处理函数

	router.GET("/search", controller.Search)
	router.GET("/chart", controller.Chart)
	router.GET("/hot_stock", controller.HotStock)

	// 获取扫码场景二维码
	router.GET("/scene_qr_code", controller.SceneQRCode)

	router.GET("/wx/check", controller.WxCheck)
	//router.POST("/wx/check", controller.WxCheck)

	// 微信公众号 关注 & 取消关注 callback
	router.POST("/wx/check", controller.Follow)
	//router.POST("/wx/unfollow", controller.Unfollow)

	// 微信支付成功 callback
	router.POST("/wx/pay", controller.Pay)

	go job.RefreshToken()

	router.Run(":9000")

	/*
		http.HandleFunc("/", HelloServer)
		http.ListenAndServe(":9000", nil)

		return

		// 获取数据
		history, err := eastmoney.NewEastMoney().QueryHistoricalFinaMainData(context.Background(), "603259.SH")
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

		//

		tableX := []interface{}{nil}
		KCFJCXSYJLR := []interface{}{"扣非净利润(亿元)"}
		KCFJCXSYJLRTZ := []interface{}{"扣非净利润同比增长(%)"}
		Totaloperatereve := []interface{}{"营业收入(亿元)"}
		Totaloperaterevetz := []interface{}{"营业收入同比增长(%)"}

		for _, item := range usefulHistory {
			tableX = append(tableX, item.ReportYear)
			KCFJCXSYJLR = append(KCFJCXSYJLR, utils.ConvertToBillions(item.Kcfjcxsyjlr))
			KCFJCXSYJLRTZ = append(KCFJCXSYJLRTZ, utils.FloatFormat(item.Kcfjcxsyjlrtz))
			Totaloperatereve = append(Totaloperatereve, utils.ConvertToBillions(item.Totaloperatereve))
			Totaloperaterevetz = append(Totaloperaterevetz, utils.FloatFormat(item.Totaloperaterevetz))
		}

		_, err = f.NewSheet("关键指标")
		if err != nil {
			return
		}

		for idx, row := range [][]interface{}{
			tableX,
			KCFJCXSYJLR,
			KCFJCXSYJLRTZ,
			Totaloperatereve,
			Totaloperaterevetz,
		} {
			cell, err := excelize.CoordinatesToCellName(1, idx+1)
			if err != nil {
				fmt.Println(err)
				return
			}
			if err := f.SetSheetRow("关键指标", cell, &row); err != nil {
				fmt.Println(err)
				return
			}
		}
		columnLetter := byte(len(tableX) + 64)

		fmt.Println(columnLetter)

		_, err = f.NewSheet("可视化报告")
		if err != nil {
			return
		}

		if err := f.SetRowVisible("可视化报告", 0, true); err != nil {
			fmt.Println(err)
		}

		if err := f.AddChart("可视化报告", "E1", &excelize.Chart{
			Type: excelize.Col,
			Series: []excelize.ChartSeries{
				{
					Name:       "关键指标!$A$2",
					Categories: fmt.Sprintf("关键指标!$%c$1:$B$1", columnLetter),
					Values:     fmt.Sprintf("关键指标!$%c$2:$B$2", columnLetter),
				},
			},
			Format: excelize.GraphicOptions{
				OffsetX: 15,
				OffsetY: 10,
			},
			Legend: excelize.ChartLegend{
				Position: "bottom_left",
			},
			Title: []excelize.RichTextRun{
				{
					Text: "扣非净利润及其增长率",
				},
			},
			PlotArea: excelize.ChartPlotArea{
				ShowCatName:     false,
				ShowLeaderLines: false,
				ShowPercent:     true,
				ShowSerName:     false,
				ShowVal:         true,
			},
			ShowBlanksAs: "zero",
			Dimension: excelize.ChartDimension{
				Width:  480 * 1.5,
				Height: 260 * 1.5,
			},
		}, &excelize.Chart{
			Type: excelize.Line,
			Series: []excelize.ChartSeries{
				{
					Name:       "关键指标!$A$3",
					Categories: fmt.Sprintf("关键指标!$%c$1:$B$1", columnLetter),
					Values:     fmt.Sprintf("关键指标!$%c$3:$B$3", columnLetter),
					Marker: excelize.ChartMarker{
						Symbol: "none", Size: 10,
					},
				},
			},
			Format: excelize.GraphicOptions{
				OffsetX: 15,
				OffsetY: 10,
			},
			Title: []excelize.RichTextRun{
				{
					Text: "扣非净利润同比增长",
				},
			},
			PlotArea: excelize.ChartPlotArea{
				ShowCatName:     false,
				ShowLeaderLines: false,
				ShowPercent:     true,
				ShowSerName:     false,
				ShowVal:         false,
			},
		}); err != nil {
			fmt.Println(err)
			return
		}

		if err := f.AddChart("可视化报告", "E26", &excelize.Chart{
			Type: excelize.Col,
			Series: []excelize.ChartSeries{
				{
					Name:       "关键指标!$A$4",
					Categories: fmt.Sprintf("关键指标!$%c$1:$B$1", columnLetter),
					Values:     fmt.Sprintf("关键指标!$%c$4:$B$4", columnLetter),
				},
			},
			Format: excelize.GraphicOptions{
				OffsetX: 15,
				OffsetY: 10,
			},
			Legend: excelize.ChartLegend{
				Position: "bottom_left",
			},
			Title: []excelize.RichTextRun{
				{
					Text: "营业收入及其增长率",
				},
			},
			PlotArea: excelize.ChartPlotArea{
				ShowCatName:     false,
				ShowLeaderLines: false,
				ShowPercent:     true,
				ShowSerName:     false,
				ShowVal:         true,
			},
			ShowBlanksAs: "zero",
			Dimension: excelize.ChartDimension{
				Width:  480 * 1.5,
				Height: 260 * 1.5,
			},
		}, &excelize.Chart{
			Type: excelize.Line,
			Series: []excelize.ChartSeries{
				{
					Name:       "关键指标!$A$5",
					Categories: fmt.Sprintf("关键指标!$%c$1:$B$1", columnLetter),
					Values:     fmt.Sprintf("关键指标!$%c$5:$B$5", columnLetter),
					Marker: excelize.ChartMarker{
						Symbol: "none", Size: 10,
					},
				},
			},
			PlotArea: excelize.ChartPlotArea{
				ShowCatName:     false,
				ShowLeaderLines: false,
				ShowPercent:     true,
				ShowSerName:     false,
				ShowVal:         false,
			},
			YAxis: excelize.ChartAxis{MajorGridLines: true},
		}); err != nil {
			fmt.Println(err)
			return
		}

		// 保存工作簿
		if err := f.SaveAs("Book1.xlsx"); err != nil {
			fmt.Println(err)
		}

	*/
}
