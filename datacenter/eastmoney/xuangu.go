package eastmoney

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"time"
)

type XuanguData struct {
	SECUCODE         string  `json:"SECUCODE"`
	SECURITYCODE     string  `json:"SECURITY_CODE"`
	SECURITYNAMEABBR string  `json:"SECURITY_NAME_ABBR"`
	NEWPRICE         float64 `json:"NEW_PRICE"`
	CHANGERATE       float64 `json:"CHANGE_RATE"`
	VOLUMERATIO      float64 `json:"VOLUME_RATIO"`
	HIGHPRICE        float64 `json:"HIGH_PRICE"`
	LOWPRICE         float64 `json:"LOW_PRICE"`
	PRECLOSEPRICE    float64 `json:"PRE_CLOSE_PRICE"`
	VOLUME           int     `json:"VOLUME"`
	DEALAMOUNT       float64 `json:"DEAL_AMOUNT"`
	TURNOVERRATE     float64 `json:"TURNOVERRATE"`
	TOTALMARKETCAP   int64   `json:"TOTAL_MARKET_CAP"`
	POPULARITYRANK   int     `json:"POPULARITY_RANK"`
	MAXTRADEDATE     string  `json:"MAX_TRADE_DATE"`
}

type RespXuanGuData struct {
	Version interface{} `json:"version"`
	Result  struct {
		Nextpage    bool         `json:"nextpage"`
		Currentpage int          `json:"currentpage"`
		Data        []XuanguData `json:"data"`
		Count       int          `json:"count"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Url     string `json:"url"`
}

// QueryFinaGincomeData 获取财务分析利润表数据，最新数据在最前面
func (e EastMoney) Xuangu(ctx context.Context) ([]XuanguData, error) {
	apiurl := "https://data.eastmoney.com/dataapi/xuangu/list"
	params := map[string]string{
		"source": "SELECT_SECURITIES",
		"client": "APP",
		"sty":    "SECUCODE,SECURITY_CODE,SECURITY_NAME_ABBR,NEW_PRICE,CHANGE_RATE,VOLUME_RATIO,HIGH_PRICE,LOW_PRICE,PRE_CLOSE_PRICE,VOLUME,DEAL_AMOUNT,TURNOVERRATE,TOTAL_MARKET_CAP,POPULARITY_RANK",
		"filter": "(TOTAL_MARKET_CAP>50000000000)(POPULARITY_RANK>0)(POPULARITY_RANK<=500)",
		"ps":     "50",
		"sr":     "-1",
		"st":     "CHANGE_RATE",
	}
	logrus.Debug(ctx, "EastMoney QueryFinaGincomeData "+apiurl+" begin", zap.Any("params", params))
	beginTime := time.Now()
	fmt.Println(params)
	resp, err := e.HTTPClient.R().SetQueryParams(params).Get(apiurl)
	latency := time.Now().Sub(beginTime).Milliseconds()
	logrus.Debug(
		ctx,
		"EastMoney QueryFinaGincomeData "+apiurl+" end",
		zap.Int64("latency(ms)", latency),
		// zap.Any("resp", resp),
	)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("%#v", resp)
	}

	respXuanGuData := &RespXuanGuData{}
	_ = json.Unmarshal(resp.Body(), respXuanGuData)

	return respXuanGuData.Result.Data, nil
}
