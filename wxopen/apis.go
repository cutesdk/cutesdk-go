package wxopen

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// FetchComponentAccessToken: request api to fetch component_access_token
func (c *Client) FetchComponentAccessToken() (request.Result, error) {
	componentVerifyTicket, err := c.GetComponentVerifyTicket()
	if err != nil {
		return nil, fmt.Errorf("get component_verify_ticket failed: %v", err)
	}

	uri := "/cgi-bin/component/api_component_token"

	res, err := c.Post(uri, map[string]interface{}{
		"component_verify_ticket": componentVerifyTicket,
		"component_appid":         c.GetComponentAppid(),
		"component_appsecret":     c.GetComponentAppsecret(),
	})

	return res, err
}
