package wxmp

import (
	"fmt"

	"github.com/idoubi/goutils"
	"github.com/idoubi/goutils/crypt"
)

// JssdkConfig: js-sdk config
type JssdkConfig struct {
	Debug     bool     `json:"debug"`
	Appid     string   `json:"appId"`
	Timestamp string   `json:"timestamp"`
	NonceStr  string   `json:"nonceStr"`
	Signature string   `json:"signature"`
	JsApiList []string `json:"jsApiList"`
}

// GetJssdkConfig: get js-sdk config
func (cli *Client) GetJssdkConfig(url string) (*JssdkConfig, error) {
	jsapiTicket, err := cli.GetJsapiTicket()
	if err != nil {
		return nil, fmt.Errorf("get jsapi_ticket failed: %v", err)
	}

	nonceStr := goutils.NonceStr(16)
	timestamp := goutils.TimestampStr()

	signStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsapiTicket, nonceStr, timestamp, url)

	signature := crypt.Sha1Encode([]byte(signStr))

	jssdkConfig := &JssdkConfig{
		Debug:     false,
		Appid:     cli.GetAppid(),
		Timestamp: timestamp,
		NonceStr:  nonceStr,
		Signature: signature,
		JsApiList: []string{},
	}

	return jssdkConfig, nil
}
