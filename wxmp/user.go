package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// GetUserList: get user list
func (c *Client) GetUserList(nextOpenid string) (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/user/get?access_token=%s", accessToken)

	params := map[string]interface{}{
		"next_openid": nextOpenid,
	}

	res, err := c.Get(uri, params)

	return res, err
}

// GetUserInfo: get userinfo
func (c *Client) GetUserInfo(openid string) (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/user/info?access_token=%s", accessToken)

	params := map[string]interface{}{
		"openid": openid,
		"lang":   "zh_CN",
	}

	res, err := c.Get(uri, params)

	return res, err
}
