package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/tidwall/sjson"
)

// CreateQrcode: create qrcode
func (c *Client) CreateQrcode(params map[string]interface{}) (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/qrcode/create?access_token=%s", accessToken)

	res, err := c.Post(uri, params)

	if err != nil {
		return nil, err
	}

	ticket := res.Get("ticket").String()
	if ticket != "" {
		longUrl := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s", ticket)
		if newRes, err := sjson.Set(res.String(), "long_url", longUrl); err == nil {
			return request.Result(newRes), nil
		}
	}

	return res, err
}
