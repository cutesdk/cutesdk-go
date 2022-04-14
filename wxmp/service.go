package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// SendServiceMsg: send service message
func (c *Client) SendServiceMsg(openid, msgtype string, params map[string]interface{}) (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/message/custom/send?access_token=%s", accessToken)

	data := map[string]interface{}{
		"touser":  openid,
		"msgtype": msgtype,
		msgtype:   params,
	}

	res, err := c.Post(uri, data)

	return res, err
}
