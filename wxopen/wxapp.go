package wxopen

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// RegisterPersonalWxapp: register personal wxapp
func (c *Client) RegisterPersonalWxapp(name, wechat string, args ...string) (request.Result, error) {
	componentAccessToken, err := c.GetComponentAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get component_access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/wxa/component/fastregisterpersonalweapp?action=create&component_access_token=%s", componentAccessToken)

	params := map[string]interface{}{
		"idname": name,
		"wxuser": wechat,
	}
	if len(args) > 0 {
		params["component_phone"] = args[0]
	}

	res, err := c.Post(uri, params)

	return res, err
}
