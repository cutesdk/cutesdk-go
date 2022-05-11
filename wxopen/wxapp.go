package wxopen

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// RegisterCompanyWxapp: register personal wxapp
func (c *Client) RegisterCompanyWxapp(params map[string]interface{}) (request.Result, error) {
	componentAccessToken, err := c.GetComponentAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get component_access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/component/fastregisterweapp?action=create&component_access_token=%s", componentAccessToken)

	res, err := c.Post(uri, params)

	return res, err
}

// RegisterPersonalWxapp: register personal wxapp
func (c *Client) RegisterPersonalWxapp(name, wechat, phone string) (request.Result, error) {
	componentAccessToken, err := c.GetComponentAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get component_access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/wxa/component/fastregisterpersonalweapp?action=create&component_access_token=%s", componentAccessToken)

	params := map[string]interface{}{
		"idname":          name,
		"wxuser":          wechat,
		"component_phone": phone,
	}

	res, err := c.Post(uri, params)

	return res, err
}
