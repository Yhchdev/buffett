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

	// 存货周转率
	chzzls := []interface{}{}
	// 总资产周转率
	toazzl := []interface{}{}

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

		chzzls = append(chzzls, utils.FloatFormat(item.Chzzl))
		toazzl = append(toazzl, utils.FloatFormat(item.Toazzl))
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
		portraits = append(portraits, utils.CorporatePortrait(item.NetcashOperate, item.NetcashInvest, item.NetcashFinance))
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

	totalcurrentassets := []interface{}{}
	totalcurrentliabs := []interface{}{}

	// 长期资本负载率
	longTermDebtRatios := []interface{}{}
	// 资产负载率
	assetsLiabilities := []interface{}{}
	// 有息负债率
	interestBearing := []interface{}{}

	// 有息负债
	youxiLiabilities := []interface{}{}

	// 货币资金
	monetaryfunds := []interface{}{}
	// 平均流动资产
	averageTotalcurrentassets := []interface{}{}
	// 平均固定资产
	averageFixedasset := []interface{}{}

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

		totalcurrentassets = append(totalcurrentassets, item.TOTALCURRENTASSETS)
		totalcurrentliabs = append(totalcurrentliabs, item.TOTALCURRENTLIAB)

		longTermDebtRatios = append(longTermDebtRatios, utils.FloatFormat(item.TOTALNONCURRENTLIAB/(item.TOTALNONCURRENTLIAB+item.TOTALEQUITY)))

		assetsLiabilities = append(assetsLiabilities, utils.FloatFormat(item.TOTALLIABILITIES/item.TOTALASSETS))

		interestBearing = append(interestBearing, utils.FloatFormat(item.NONCURRENTLIAB1YEAR/item.TOTALLIABILITIES))
		youxiLiabilities = append(youxiLiabilities, utils.ConvertToBillions(item.NONCURRENTLIAB1YEAR))

		monetaryfunds = append(monetaryfunds, utils.ConvertToBillions(item.MONETARYFUNDS))

		if i == 0 {
			averageTotalcurrentassets = append(averageTotalcurrentassets, item.TOTALCURRENTASSETS)
			averageFixedasset = append(averageFixedasset, item.FIXEDASSET)
		} else {
			averageTotalcurrentassets = append(averageTotalcurrentassets, (item.TOTALCURRENTASSETS+balance[i-1].TOTALCURRENTASSETS)/2)
			averageFixedasset = append(averageFixedasset, (item.FIXEDASSET+balance[i-1].FIXEDASSET)/2)
		}
	}

	//  资本收益率
	zibenshouyis := []interface{}{}
	// 流动比率
	liudongs := []interface{}{}

	// 现金比率
	cashcfuzais := []interface{}{}

	// 流动资产周转率
	currentAssetTurnoverRatio := []interface{}{}

	// 固定资产周转率
	fixedAssetTurnoverRatio := []interface{}{}

	for i := 0; i < len(totalparentequity); i++ {
		zibenshouyis = append(zibenshouyis, utils.FloatFormat(cast.ToFloat64(KCFJCXSYJLR[i])/cast.ToFloat64(totalparentequity[i])))
		liudongs = append(liudongs, utils.FloatFormat(cast.ToFloat64(totalcurrentassets[i])/cast.ToFloat64(totalcurrentliabs[i])))
		cashcfuzais = append(cashcfuzais, utils.FloatFormat(cast.ToFloat64(MONETARYFUNDS[i])/utils.ConvertToBillions(cast.ToFloat64(totalcurrentliabs[i]))))
		currentAssetTurnoverRatio = append(currentAssetTurnoverRatio, utils.FloatFormat(cashFlow[i].SalesServices/cast.ToFloat64(averageTotalcurrentassets[i])))
		fixedAssetTurnoverRatio = append(fixedAssetTurnoverRatio, utils.FloatFormat(cashFlow[i].SalesServices/cast.ToFloat64(averageFixedasset[i])))
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
				Name:  "营业收入",
				Type:  "bar",
				Y:     Totaloperatereve,
				YType: "value",
			},
			{
				Name:  "增长率",
				Type:  "line",
				Y:     Totaloperaterevetz,
				YType: "value",
			},
		},
	}, Char{
		Name: "盈利增长能力",
		Series: []Series{
			{
				Name:  "营业收入增长率",
				Type:  "line",
				Y:     Totaloperaterevetz,
				YType: "value",
			},
			{
				Name:  "净利润增长率",
				Type:  "line",
				Y:     Parentnetprofittz,
				YType: "value",
			},
			{
				Name:  "扣非净利润增长率",
				Type:  "line",
				Y:     KCFJCXSYJLRTZ,
				YType: "value",
			},
		},
	}, Char{
		Name: "盈利能力",
		Series: []Series{
			{
				Name:  "毛利率",
				Type:  "line",
				Y:     Xsmll,
				YType: "value",
			},
			{
				Name:  "净利率",
				Type:  "line",
				Y:     Xsjll,
				YType: "value",
			},
		},
	}, Char{
		Name: "核心净利润及其贡献率",
		Series: []Series{
			{
				Name:  "核心净利润",
				Type:  "bar",
				Y:     coreProfit,
				YType: "value",
			},
			{
				Name:  "核心净利润率",
				Type:  "line",
				Y:     coreProfitCompareOperateIncome,
				YType: "value",
			},
		},
	}, Char{
		Name: "净利润与营收现金净流量",
		Series: []Series{
			{
				Name:  "经营活动净流量",
				Type:  "line",
				Y:     netcashOperate,
				YType: "value",
			},
			{
				Name:  "净利润",
				Type:  "line",
				Y:     Parentnetprofit,
				YType: "value",
			},
			{
				Name:  "扣非净利润",
				Type:  "line",
				Y:     KCFJCXSYJLR,
				YType: "value",
			},
		},
	}, Char{
		Name: "业绩真实性分析",
		Series: []Series{
			{
				Name:  "净现比",
				Type:  "line",
				Y:     netcashOperateCompareKCFJCXSYJLR,
				YType: "value",
			},
			{
				Name:  "核现比",
				Type:  "line",
				Y:     coreProfitCompareKCFJCXSYJLR,
				YType: "value",
			},
			{
				Name:  "收现比",
				Type:  "line",
				Y:     salesServicesCompareTotaloperatereve,
				YType: "value",
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
				YType:    "value",
			},
			{
				Name:     "货币资金",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        MONETARYFUNDS,
				YType:    "value",
			},
			{
				Name:     "存货",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        INVENTORY,
				YType:    "value",
			},
			{
				Name:     "在建工程",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        CIP,
				YType:    "value",
			},
			{
				Name:     "固定资产",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        FIXEDASSET,
				YType:    "value",
			},
			{
				Name:     "应收账款",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        NOTEACCOUNTSRECE,
				YType:    "value",
			},
		},
	}, Char{
		Name: "经营性资产及其占比",
		Series: []Series{
			{
				Name:  "经营性资产",
				Type:  "bar",
				Y:     jingyingxingzichans,
				YType: "value",
			},
			{
				Name:  "经营性资产占比",
				Type:  "line",
				Y:     jingyingxingzichanProportion,
				YType: "value",
			},
		},
	}, Char{
		Name: "固定资产率",
		Series: []Series{
			{
				Name:  "固定资产率",
				Type:  "line",
				Y:     fixedassetProportion,
				YType: "value",
			},
		},
	}, Char{
		Name: "预付款及其占营收的比例",
		Series: []Series{
			{
				Name:  "预付款",
				Type:  "bar",
				Y:     prepayments,
				YType: "value",
			},
			{
				Name:  "预付款占营收比例",
				Type:  "line",
				Y:     prepaymentsProportion,
				YType: "value",
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
	}, Char{
		Name: "短期偿债能力分析",
		Series: []Series{
			{
				Name: "流动比率",
				Type: "line",
				Y:    liudongs,
			},
			{
				Name: "现金比率",
				Type: "line",
				Y:    cashcfuzais,
			},
		},
	}, Char{
		Name: "长期偿债能力分析",
		Series: []Series{
			{
				Name: "长期资本负债率",
				Type: "line",
				Y:    longTermDebtRatios,
			},
		},
	}, Char{
		Name: "现金和有息负债",
		Series: []Series{
			{
				Name:     "现金和现金等价物",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        monetaryfunds,
			},
			{
				Name:     "有利息负债",
				Type:     "bar",
				Emphasis: Emphasis{Focus: "series"},
				Stack:    "total",
				Y:        youxiLiabilities,
			},
		},
	}, Char{
		Name: "存货周转率",
		Series: []Series{
			{
				Name: "存货周转率",
				Type: "line",
				Y:    chzzls,
			},
		},
	}, Char{
		Name: "运营能力",
		Series: []Series{
			{
				Name: "流动资产周转率",
				Type: "line",
				Y:    currentAssetTurnoverRatio,
			}, {
				Name: "固定资产周转率",
				Type: "line",
				Y:    fixedAssetTurnoverRatio,
			}, {
				Name: "总资产周转率",
				Type: "line",
				Y:    toazzl,
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
