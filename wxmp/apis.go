package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/idoubi/goutils"
	"github.com/idoubi/goutils/crypt"
)

// FetchAccessToken: request get_access_token api
func (c *Client) FetchAccessToken() (request.Result, error) {
	uri := "/cgi-bin/token"

	res, err := c.Get(uri, map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      c.GetAppid(),
		"secret":     c.GetSecret(),
	})

	return res, err
}

// FetchJsapiTicket: request get_jsapi_ticket api
func (c *Client) FetchJsapiTicket() (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", accessToken)

	res, err := c.Get(uri)

	return res, err
}

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
func (c *Client) GetJssdkConfig(url string) (*JssdkConfig, error) {
	jsapiTicket, err := c.GetJsapiTicket()
	if err != nil {
		return nil, fmt.Errorf("get jsapi_ticket failed: %v", err)
	}

	nonceStr := goutils.NonceStr(16)
	timestamp := goutils.TimestampStr()

	signStr := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsapiTicket, nonceStr, timestamp, url)

	signature := crypt.Sha1Encode([]byte(signStr))

	jssdkConfig := &JssdkConfig{
		Debug:     false,
		Appid:     c.GetAppid(),
		Timestamp: timestamp,
		NonceStr:  nonceStr,
		Signature: signature,
		JsApiList: []string{},
	}

	return jssdkConfig, nil
}
