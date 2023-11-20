package controller

import (
	"fmt"
	"github.com/Yhchdev/buffett/datacenter/eastmoney"
	"github.com/Yhchdev/buffett/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"sort"
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

	if upperStr == "" && reportType == "" {
		fmt.Println("非法参数")
		return
	}

	// 获取经营分析数据
	history, err := eastmoney.NewEastMoney().QueryHistoricalFinaMainData(c, upperStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	usefulHistory := history.FilterByReportType(eastmoney.FinaReportType(reportType))

	if len(usefulHistory) > maxHistory {
		usefulHistory = usefulHistory[len(usefulHistory)-9:]
	}

	// 查询的年份和月份列表
	incomeDate := utils.GincomeReportDateParams(reportType, len(usefulHistory), cast.ToInt(usefulHistory[len(usefulHistory)-1].ReportYear))
	fmt.Println(incomeDate)

	// 获取利润表数据
	incomes, err := eastmoney.NewEastMoney().QueryFinaGincomeData(c, upperStr, incomeDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Slice(incomes, func(i, j int) bool {
		return incomes[i].ReportDate < incomes[j].ReportDate
	})

	// 获取现金流量表数据

	cashFlow, err := eastmoney.NewEastMoney().QueryFinaCashflowData(c, upperStr, incomeDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Slice(cashFlow, func(i, j int) bool {
		return cashFlow[i].ReportDate < cashFlow[j].ReportDate
	})

	//fmt.Println(incomes)

	tableX := []interface{}{}
	KCFJCXSYJLR := []interface{}{}
	KCFJCXSYJLRTZ := []interface{}{}
	Totaloperatereve := []interface{}{}
	Totaloperaterevetz := []interface{}{}
	Parentnetprofit := []interface{}{}
	Parentnetprofittz := []interface{}{}
	Xsmll := []interface{}{}
	Xsjll := []interface{}{}

	for _, item := range usefulHistory {
		tableX = append(tableX, item.ReportYear)
		KCFJCXSYJLR = append(KCFJCXSYJLR, utils.ConvertToBillions(item.Kcfjcxsyjlr))
		KCFJCXSYJLRTZ = append(KCFJCXSYJLRTZ, utils.FloatFormat(item.Kcfjcxsyjlrtz))
		Totaloperatereve = append(Totaloperatereve, utils.ConvertToBillions(item.Totaloperatereve))
		Totaloperaterevetz = append(Totaloperaterevetz, utils.FloatFormat(item.Totaloperaterevetz))
		Parentnetprofittz = append(Parentnetprofittz, utils.FloatFormat(item.Parentnetprofittz))
		Xsmll = append(Xsmll, utils.FloatFormat(item.Xsmll))
		Xsjll = append(Xsjll, utils.FloatFormat(item.Xsjll))
		Parentnetprofit = append(Parentnetprofit, utils.ConvertToBillions(item.Parentnetprofit))
	}

	// 核心利润
	coreProfit := []interface{}{}
	coreProfitCompareOperateIncome := []interface{}{}

	for _, item := range incomes {
		coreP := utils.ConvertToBillions(item.OperateIncome - item.OperateTaxAdd - item.OperateCost - item.ManageExpense -
			item.SaleExpense - item.FinanceExpense)
		coreProfit = append(coreProfit, coreP)

		coreProfitCIncome := utils.FloatFormat(coreP/utils.ConvertToBillions(item.OperateIncome)) * 100

		coreProfitCompareOperateIncome = append(coreProfitCompareOperateIncome, coreProfitCIncome)
	}

	netcashOperate := []interface{}{}
	salesServices := []interface{}{}

	for _, item := range cashFlow {
		netcashOperate = append(netcashOperate, utils.ConvertToBillions(item.NetcashOperate))
		salesServices = append(salesServices, utils.ConvertToBillions(item.SalesServices))
	}

	// 净现比
	netcashOperateCompareKCFJCXSYJLR := []interface{}{}
	// 核现比
	coreProfitCompareKCFJCXSYJLR := []interface{}{}
	// 销售收现比
	salesServicesCompareTotaloperatereve := []interface{}{}

	//  比较
	for i := 0; i < len(usefulHistory); i++ {
		netcashOperateCompareKCFJCXSYJLR = append(netcashOperateCompareKCFJCXSYJLR, utils.FloatFormat(cast.ToFloat64(netcashOperate[i])/cast.ToFloat64(KCFJCXSYJLR[i])))
		coreProfitCompareKCFJCXSYJLR = append(coreProfitCompareKCFJCXSYJLR, utils.FloatFormat(cast.ToFloat64(coreProfit[i])/cast.ToFloat64(KCFJCXSYJLR[i])))
		salesServicesCompareTotaloperatereve = append(salesServicesCompareTotaloperatereve, utils.FloatFormat(cast.ToFloat64(salesServices[i])/cast.ToFloat64(Totaloperatereve[i])))
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
	}, Char{
		Name: "盈利增长能力",
		Series: []Series{
			{
				Name: "营业收入增长率",
				Type: "line",
				X:    tableX,
				Y:    Totaloperaterevetz,
			},
			{
				Name: "净利润增长率",
				Type: "line",
				X:    tableX,
				Y:    Parentnetprofittz,
			},
			{
				Name: "扣非净利润增长率",
				Type: "line",
				X:    tableX,
				Y:    KCFJCXSYJLRTZ,
			},
		},
	}, Char{
		Name: "盈利能力",
		Series: []Series{
			{
				Name: "毛利率",
				Type: "line",
				X:    tableX,
				Y:    Xsmll,
			},
			{
				Name: "净利率",
				Type: "line",
				X:    tableX,
				Y:    Xsjll,
			},
		},
	}, Char{
		Name: "核心净利润及其贡献率",
		Series: []Series{
			{
				Name: "核心净利润",
				Type: "bar",
				X:    tableX,
				Y:    coreProfit,
			},
			{
				Name: "核心净利润率",
				Type: "line",
				X:    tableX,
				Y:    coreProfitCompareOperateIncome,
			},
		},
	}, Char{
		Name: "净利润与营收现金净流量",
		Series: []Series{
			{
				Name: "经营活动净流量",
				Type: "line",
				X:    tableX,
				Y:    netcashOperate,
			},
			{
				Name: "净利润",
				Type: "line",
				X:    tableX,
				Y:    Parentnetprofit,
			},
			{
				Name: "扣非净利润",
				Type: "line",
				X:    tableX,
				Y:    KCFJCXSYJLR,
			},
		},
	}, Char{
		Name: "业绩真实性分析",
		Series: []Series{
			{
				Name: "净现比",
				Type: "line",
				X:    tableX,
				Y:    netcashOperateCompareKCFJCXSYJLR,
			},
			{
				Name: "核现比",
				Type: "line",
				X:    tableX,
				Y:    coreProfitCompareKCFJCXSYJLR,
			},
			{
				Name: "收现比",
				Type: "line",
				X:    tableX,
				Y:    salesServicesCompareTotaloperatereve,
			},
		},
	})

	// 处理请求并返回响应
	c.JSON(200, gin.H{
		"charts": charts,
	})

}
