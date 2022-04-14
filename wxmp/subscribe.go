package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// SendSubscribeMsg: send subscribe msg
func (c *Client) SendSubscribeMsg(tplId, openid string, data map[string]interface{}, page string, miniprogram map[string]interface{}) (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/message/subscribe/bizsend?access_token=%s", accessToken)

	params := map[string]interface{}{
		"touser":      openid,
		"template_id": tplId,
		"page":        page,
		"data":        data,
		"miniprogram": miniprogram,
	}

	res, err := c.Post(uri, params)

	return res, err
}
