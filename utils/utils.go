package utils

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"strings"
)

func ConvertToBillions(n float64) float64 {
	return math.Round(n/1e8*100) / 100
}

func FloatFormat(n float64) float64 {
	return math.Round(n*100) / 100
}

func FloatFormatToGrowth(n float64) float64 {
	return (math.Round(n*100) / 100) * 100
}

func GincomeReportDateParams(reportType string, nums int, startYear int) string {

	suffix := ""

	switch reportType {
	case "年报":
		suffix = "12-31"
	case "三季报":
		suffix = "09-30"
	case "中报":
		suffix = "06-30"
	case "一季报":
		suffix = "03-31"
	}

	params := make([]string, 0)
	for i := 0; i < nums; i++ {

		if i > 0 {
			startYear -= 1
		}

		params = append(params, fmt.Sprintf("'%s-%s'", cast.ToString(startYear), suffix))
	}

	return strings.Join(params, ",")
}
