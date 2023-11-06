package main

import (
	"context"
	"fmt"
	"github.com/Yhchdev/buffett/utils"
	"github.com/xuri/excelize/v2"

	"github.com/Yhchdev/buffett/datacenter/eastmoney"
)

func main() {

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
	KCFJCXSYJLR := []interface{}{"扣非净利润"}
	KCFJCXSYJLRTZ := []interface{}{"扣非净利润同比增长（%）"}

	for _, item := range usefulHistory {
		tableX = append(tableX, item.ReportYear)
		KCFJCXSYJLR = append(KCFJCXSYJLR, utils.ConvertToBillions(item.Kcfjcxsyjlr))
		KCFJCXSYJLRTZ = append(KCFJCXSYJLRTZ, item.Kcfjcxsyjlrtz)
	}

	for idx, row := range [][]interface{}{
		tableX,
		KCFJCXSYJLR,
		KCFJCXSYJLRTZ,
	} {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := f.SetSheetRow("Sheet1", cell, &row); err != nil {
			fmt.Println(err)
			return
		}
	}
	columnLetter := byte(len(tableX) + 64)

	fmt.Println(columnLetter)

	if err := f.AddChart("Sheet1", "E1", &excelize.Chart{
		Type: excelize.Col,
		Series: []excelize.ChartSeries{
			{
				Name:       "Sheet1!$A$2",
				Categories: fmt.Sprintf("Sheet1!$%c$1:$B$1", columnLetter),
				Values:     fmt.Sprintf("Sheet1!$%c$2:$B$2", columnLetter),
			},
		},
		Format: excelize.GraphicOptions{
			OffsetX: 15,
			OffsetY: 10,
		},
		Legend: excelize.ChartLegend{
			Position: "left",
		},
		Title: []excelize.RichTextRun{
			{
				Text: "扣非净利润",
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
	}, &excelize.Chart{
		Type: excelize.Line,
		Series: []excelize.ChartSeries{
			{
				Name:       "扣非净利润同比增长（%）",
				Categories: fmt.Sprintf("Sheet1!$%c$1:$B$1", columnLetter),
				Values:     fmt.Sprintf("Sheet1!$%c$3:$B$3", columnLetter),
				Marker: excelize.ChartMarker{
					Symbol: "none", Size: 10,
				},
			},
		},
		Format: excelize.GraphicOptions{
			OffsetX: 15,
			OffsetY: 10,
		},
		Legend: excelize.ChartLegend{
			Position: "right",
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

	// 保存工作簿
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
