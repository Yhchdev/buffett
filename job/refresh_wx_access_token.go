package job

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"time"
)

// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html

type WxToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

var AccessToken string

func RefreshToken() {

	HTTPClient := resty.New()

	for {
		params := map[string]string{
			"grant_type": "client_credential",
			"appid":      "wx1987d819f8dc00ca",
			"secret":     "3c04de17597a4e1e958135882405a71e",
		}

		resp, err := HTTPClient.R().SetQueryParams(params).Get("https://api.weixin.qq.com/cgi-bin/token")
		if err != nil {
			logrus.Error(err)
			continue
		}

		wxToken := WxToken{}
		if err := json.Unmarshal(resp.Body(), &wxToken); err != nil {
			logrus.Error(err)
			continue
		}

		AccessToken = wxToken.AccessToken
		logrus.Println("get token success:", AccessToken)

		time.Sleep(time.Second * 7000)
	}
}
