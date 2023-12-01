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

type CharsResp struct {
	Charts []Char        `json:"charts"`
	X      []interface{} `json:"x"` // 所有图的x是一致的
}

type Char struct {
	Name   string   `json:"name"`
	Series []Series `json:"series"`
}

type Series struct {
	Type     string   `json:"type"`     // 图的类型  line bar
	Name     string   `json:"name"`     // 图的名称
	Stack    string   `json:"stack"`    // 柱状图堆叠
	Emphasis Emphasis `json:"emphasis"` // 高亮
	//MarkPoint []interface{} `json:"markPoint"`
	Y     []interface{} `json:"y"`
	YType string        `json:"YType"`
}

type Emphasis struct {
	Focus string `json:"focus" default:"series"`
}

//type MarkPoint s

func Chart(c *gin.Context) {

	upperStr := strings.ToUpper(c.Query("secucode"))
	reportType := c.Query("report")

	if upperStr == "" || reportType == "" {
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

	// 获取资产负债表数据

	balance, err := eastmoney.NewEastMoney().QueryBlanceData(c, upperStr, incomeDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	sort.Slice(balance, func(i, j int) bool {
		return balance[i].REPORTDATE < balance[j].REPORTDATE
	})

	tableX := []interface{}{}
	KCFJCXSYJLR := []interface{}{}
	KCFJCXSYJLRTZ := []interface{}{}
	Totaloperatereve := []interface{}{}
	Totaloperaterevetz := []interface{}{}
	Parentnetprofit := []interface{}{}
	Parentnetprofittz := []interface{}{}
	Xsmll := []interface{}{}
	Xsjll := []interface{}{}
	roekcjqs := []interface{}{}
	zzcjlls := []interface{}{}
	roics := []interface{}{}

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

		roekcjqs = append(roekcjqs, utils.FloatFormat(item.Roekcjq))
		zzcjlls = append(zzcjlls, utils.FloatFormat(item.Zzcjll))
		roics = append(roics, utils.FloatFormat(item.Roic))
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
	portraits := []interface{}{}
	netcashInvest := []interface{}{}
	netcashFinance := []interface{}{}
	totalInvestInflowProportions := []interface{}{}
	totalOperateInflowProportions := []interface{}{}
	totalFinanceInflowProportions := []interface{}{}
	buyServicesCompareSales := []interface{}{}

	for _, item := range cashFlow {
		netcashOperate = append(netcashOperate, utils.ConvertToBillions(item.NetcashOperate))
		netcashInvest = append(netcashInvest, utils.ConvertToBillions(item.NetcashInvest))
		netcashFinance = append(netcashFinance, utils.ConvertToBillions(item.NetcashFinance))
		salesServices = append(salesServices, utils.ConvertToBillions(item.SalesServices))
		totalInflow := item.TotalInvestInflow + item.TotalOperateInflow + item.TotalFinanceInflow
		totalInvestInflowProportions = append(totalInvestInflowProportions, utils.FloatFormat(item.TotalInvestInflow/totalInflow))
		totalOperateInflowProportions = append(totalOperateInflowProportions, utils.FloatFormat(item.TotalOperateInflow/totalInflow))
		totalFinanceInflowProportions = append(totalFinanceInflowProportions, utils.FloatFormat(item.TotalFinanceInflow/totalInflow))
		buyServicesCompareSales = append(buyServicesCompareSales, utils.FloatFormat(item.BuyServices/item.SalesServices))
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

	LONGEQUITYINVEST := []interface{}{}
	MONETARYFUNDS := []interface{}{}
	INVENTORY := []interface{}{}
	CIP := []interface{}{}
	FIXEDASSET := []interface{}{}
	NOTEACCOUNTSRECE := []interface{}{}
	//PREPAYMENT := []interface{}{}
	// 经营性资产
	jingyingxingzichans := []interface{}{}

	jingyingxingzichanProportion := []interface{}{}
	fixedassetProportion := []interface{}{}
	prepayments := []interface{}{}
	prepaymentsProportion := []interface{}{}

	totalparentequity := []interface{}{}

	for i, item := range balance {
		LONGEQUITYINVEST = append(LONGEQUITYINVEST, utils.ConvertToBillions(cast.ToFloat64(item.LONGEQUITYINVEST)))
		MONETARYFUNDS = append(MONETARYFUNDS, utils.ConvertToBillions(item.MONETARYFUNDS))
		INVENTORY = append(INVENTORY, utils.ConvertToBillions(item.INVENTORY))
		CIP = append(CIP, utils.ConvertToBillions(item.CIP))
		FIXEDASSET = append(FIXEDASSET, utils.ConvertToBillions(item.FIXEDASSET))
		NOTEACCOUNTSRECE = append(NOTEACCOUNTSRECE, utils.ConvertToBillions(item.NOTEACCOUNTSRECE))
		jingyingxingzichan := item.INVENTORY + item.NOTEACCOUNTSRECE + item.PREPAYMENT + item.MONETARYFUNDS
		jingyingxingzichans = append(jingyingxingzichans, utils.ConvertToBillions(jingyingxingzichan))

		jingyingxingzichanProportion = append(jingyingxingzichanProportion, utils.FloatFormat(jingyingxingzichan/item.TOTALASSETS))
		fixedassetProportion = append(fixedassetProportion, utils.FloatFormat(item.FIXEDASSET/item.TOTALASSETS))
		prepayments = append(prepayments, utils.ConvertToBillions(item.PREPAYMENT))
		prepaymentsProportion = append(prepaymentsProportion, utils.FloatFormat(utils.ConvertToBillions(item.PREPAYMENT)/cast.ToFloat64(Totaloperatereve[i])))
		totalparentequity = append(totalparentequity, utils.ConvertToBillions(item.TOTALPARENTEQUITY))
	}

	//  资本收益率
	zibenshouyis := []interface{}{}
	for i := 0; i < len(totalparentequity); i++ {
		zibenshouyis = append(zibenshouyis, utils.FloatFormat(cast.ToFloat64(KCFJCXSYJLR[i])/cast.ToFloat64(totalparentequity[i])))
	}

	//fmt.Println("MONETARYFUNDS", MONETARYFUNDS)

	charts := make([]Char, 0)

	charts = append(charts, Char{
		Name: "扣非净利润及其增长率",
		Series: []Series{
			{
				Name: "扣非净利润",
				Type: "bar",
				Y:    KCFJCXSYJLR,
			},
			{
				Name: "增长率",
				Type: "line",
				Y:    KCFJCXSYJLRTZ,
			},
		},
	}, Char{
		Name: "营业收入及其增长率",
		Series: []Series{
			{
				Name: "营业收入",
				Type: "bar",
				Y:    Totaloperatereve,
			},
			{
				Name: "增长率",
				Type: "line",
				Y:    Totaloperaterevetz,
			},
		},
	}, Char{
		Name: "盈利增长能力",
		Series: []Series{
			{
				Name: "营业收入增长率",
				Type: "line",
				Y:    Totaloperaterevetz,
			},
			{
				Name: "净利润增长率",
				Type: "line",
				Y:    Parentnetprofittz,
			},
			{
				Name: "扣非净利润增长率",
				Type: "line",
				Y:    KCFJCXSYJLRTZ,
			},
		},
	}, Char{
		Name: "盈利能力",
		Series: []Series{
			{
				Name: "毛利率",
				Type: "line",
				Y:    Xsmll,
			},
			{
				Name: "净利率",
				Type: "line",
				Y:    Xsjll,
			},
		},
	}, Char{
		Name: "核心净利润及其贡献率",
		Series: []Series{
			{
				Name: "核心净利润",
				Type: "bar",
				Y:    coreProfit,
			},
			{
				Name: "核心净利润率",
				Type: "line",
				Y:    coreProfitCompareOperateIncome,
			},
		},
	}, Char{
		Name: "净利润与营收现金净流量",
		Series: []Series{
			{
				Name: "经营活动净流量",
				Type: "line",
				Y:    netcashOperate,
			},
			{
				Name: "净利润",
				Type: "line",
				Y:    Parentnetprofit,
			},
			{
				Name: "扣非净利润",
				Type: "line",
				Y:    KCFJCXSYJLR,
			},
		},
	}, Char{
		Name: "业绩真实性分析",
		Series: []Series{
			{
				Name: "净现比",
				Type: "line",
				Y:    netcashOperateCompareKCFJCXSYJLR,
			},
			{
				Name: "核现比",
				Type: "line",
				Y:    coreProfitCompareKCFJCXSYJLR,
			},
			{
				Name: "收现比",
				Type: "line",
				Y:    salesServicesCompareTotaloperatereve,
			},
		},
	}, Char{
		Name: "主要资产结构",
		Series: []Series{
			{
				Name:     "长期股权投资",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        LONGEQUITYINVEST,
			},
			{
				Name:     "货币资金",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        MONETARYFUNDS,
			},
			{
				Name:     "存货",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        INVENTORY,
			},
			{
				Name:     "在建工程",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        CIP,
			},
			{
				Name:     "固定资产",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        FIXEDASSET,
			},
			{
				Name:     "应收账款",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        NOTEACCOUNTSRECE,
			},
		},
	}, Char{
		Name: "经营性资产及其占比",
		Series: []Series{
			{
				Name: "经营性资产",
				Type: "bar",
				Y:    jingyingxingzichans,
			},
			{
				Name: "经营性资产占比",
				Type: "line",
				Y:    jingyingxingzichanProportion,
			},
		},
	}, Char{
		Name: "固定资产率",
		Series: []Series{
			{
				Name: "固定资产率",
				Type: "line",
				Y:    fixedassetProportion,
			},
		},
	}, Char{
		Name: "预付款及其占营收的比例",
		Series: []Series{
			{
				Name: "预付款",
				Type: "bar",
				Y:    prepayments,
			},
			{
				Name: "预付款占营收比例",
				Type: "line",
				Y:    prepaymentsProportion,
			},
		},
	}, Char{
		Name: "公司现金流画像",
		Series: []Series{
			{
				Name:  "公司类型",
				Type:  "line",
				Y:     portraits,
				YType: "category",
			},
		},
	}, Char{
		Name: "现金流结构",
		Series: []Series{
			{
				Name:     "经营活动产生的现金流净额",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        netcashOperate,
			},
			{
				Name:     "投资活动产生的现金流净额",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        netcashInvest,
			},
			{
				Name:     "筹资活动产生的现金流净额",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        netcashFinance,
			},
		},
	}, Char{
		Name: "资金流入构成分析",
		Series: []Series{
			{
				Name:     "经营活动产生的现金流入占比",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        totalOperateInflowProportions,
			},
			{
				Name:     "投资活动现金流入占比",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        totalInvestInflowProportions,
			},
			{
				Name:     "筹资活动现金流入占比",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        totalFinanceInflowProportions,
			},
		},
	}, Char{
		Name: "资金流出构成分析",
		Series: []Series{
			{
				Name:     "经营活动产生的现金流出占比",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        totalOperateInflowProportions,
			},
			{
				Name:     "投资活动现金流出占比",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        totalInvestInflowProportions,
			},
			{
				Name:     "筹资活动现金流出占比",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        totalFinanceInflowProportions,
			},
		},
	}, Char{
		Name: "现金购销比",
		Series: []Series{
			{
				Name: "现金购销比",
				Type: "line",
				Y:    buyServicesCompareSales,
			},
		},
	}, Char{
		Name: "盈利能力分析",
		Series: []Series{
			{
				Name: "ROIC(资本回报率)",
				Type: "line",
				Y:    roics,
			},
			{
				Name: "ROA(总资产收益率)",
				Type: "line",
				Y:    zzcjlls,
			},
			{
				Name: "ROE(净资产收益率)",
				Type: "line",
				Y:    roekcjqs,
			},
		},
	}, Char{
		Name: "净资产和ROE",
		Series: []Series{
			{
				Name: "ROA(总资产收益率)",
				Type: "line",
				Y:    zzcjlls,
			},
			{
				Name: "净资产",
				Type: "bar",
				Y:    totalparentequity,
			},
		},
	}, Char{
		Name: "资本收益率",
		Series: []Series{
			{
				Name: "资本收益率",
				Type: "line",
				Y:    zibenshouyis,
			},
		},
	})

	// 处理请求并返回响应
	c.JSON(200, gin.H{
		"stock_charts": CharsResp{
			Charts: charts,
			X:      tableX,
		},
	})

}
