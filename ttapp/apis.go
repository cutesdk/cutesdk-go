package ttapp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// FetchAccessToken: request get_access_token api
func (c *Client) FetchAccessToken() (request.Result, error) {
	apiPath := "/api/apps/v2/token"

	res, err := c.Post(apiPath, map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      c.GetAppid(),
		"secret":     c.GetSecret(),
	})

	return res, err
}

// Code2Session: request code2session api
func (c *Client) Code2Session(code, anonymousCode string) (request.Result, error) {
	uri := "/api/apps/v2/jscode2session"

	res, err := c.Post(uri, map[string]interface{}{
		"appid":          c.GetAppid(),
		"secret":         c.GetSecret(),
		"code":           code,
		"anonymous_code": anonymousCode,
	})

	return res, err
}

// CreateQrcode: request create_qrcode api
func (c *Client) CreateQrcode(data map[string]interface{}) (request.Result, error) {
	uri := "/api/apps/qrcode"

	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	if data == nil {
		data = map[string]interface{}{
			"access_token": accessToken,
		}
	} else {
		data["access_token"] = accessToken
	}

	res, err := c.Post(uri, data)

	return res, err
}

// SendSubscribeMsg: request send subscribe message api
func (c *Client) SendSubscribeMsg(tplId, openid string, data map[string]string, args ...string) (request.Result, error) {
	uri := "/api/apps/subscribe_notification/developer/v1/notify"

	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	page := ""
	if len(args) > 0 {
		page = args[0]
	}

	postData := map[string]interface{}{
		"access_token": accessToken,
		"app_id":       c.GetAppid(),
		"tpl_id":       tplId,
		"open_id":      openid,
		"data":         data,
		"page":         page,
	}

	res, err := c.Post(uri, postData)

	return res, err
}
