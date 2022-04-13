package wxapp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// CreateQrcode: request wxacode.createQRCode api
func (c *Client) CreateQrcode(params map[string]interface{}) (request.Result, error) {
	return c.CreateQrcodeByType(1, params)
}

// GetQrcode: request wxacode.get api
func (c *Client) GetQrcode(params map[string]interface{}) (request.Result, error) {
	return c.CreateQrcodeByType(2, params)
}

// GetUnlimitedQrcode: request wxacode.getUnlimited api
func (c *Client) GetUnlimitedQrcode(params map[string]interface{}) (request.Result, error) {
	return c.CreateQrcodeByType(3, params)
}

// CreateQrcodeByType: create qrcode request different apis
func (c *Client) CreateQrcodeByType(_type int8, params map[string]interface{}) (request.Result, error) {
	accessToken, err := c.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	var uri string

	switch _type {
	case 1:
		uri = fmt.Sprintf("/cgi-bin/wxaapp/createwxaqrcode?access_token=%s", accessToken)
	case 2:
		uri = fmt.Sprintf("/wxa/getwxacode?access_token=%s", accessToken)
	case 3:
		uri = fmt.Sprintf("/wxa/getwxacodeunlimit?access_token=%s", accessToken)
	default:
		return nil, fmt.Errorf("invalid type")
	}

	res, err := c.Post(uri, params)

	return res, err
}
