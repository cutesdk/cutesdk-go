package wxapp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
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

// Code2Session: request code2session api
func (c *Client) Code2Session(code string) (request.Result, error) {
	uri := "/sns/jscode2session"

	res, err := c.Get(uri, map[string]interface{}{
		"appid":      c.GetAppid(),
		"secret":     c.GetSecret(),
		"js_code":    code,
		"grant_type": "authorization_code",
	})

	return res, err
}

// GetPhone: get user phone
func (c *Client) GetUserPhone(code string) (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/wxa/business/getuserphonenumber?access_token=%s", accessToken)

	res, err := c.Post(uri, map[string]interface{}{
		"code": code,
	})

	return res, err
}
