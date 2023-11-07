package main

import (
	"context"
	"fmt"
	"github.com/Yhchdev/buffett/utils"
	"github.com/xuri/excelize/v2"
	"net/http"

	"github.com/Yhchdev/buffett/datacenter/eastmoney"
)

func HelloServer(resp http.ResponseWriter, r *http.Request) {

	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	// 必须，设置服务器支持的所有跨域请求的方法
	resp.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	// 服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段
	resp.Header().Set("Access-Control-Allow-Headers", "content-type")
	// 可选，设置XMLHttpRequest的响应对象能拿到的额外字段
	resp.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Headers, Token")
	// 可选，是否允许后续请求携带认证信息Cookir，该值只能是true，不需要则不设置
	resp.Header().Set("Access-Control-Allow-Credentials", "true")

	fmt.Fprintf(resp, "Hello, %s!", r.URL.Path[1:])
}

func main() {
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
}
