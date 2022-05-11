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

// StartPushTicket: start push component_verify_ticket
func (c *Client) StartPushTicket() (request.Result, error) {
	uri := "/cgi-bin/component/api_start_push_ticket"

	res, err := c.Post(uri, map[string]interface{}{
		"component_appid":  c.GetComponentAppid(),
		"component_secret": c.GetComponentAppsecret(),
	})

	return res, err
}

// CreatePreauthCode: create preauthcode
func (c *Client) CreatePreauthCode() (request.Result, error) {
	componentAccessToken, err := c.GetComponentAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get component_access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/component/api_create_preauthcode?component_access_token=%s", componentAccessToken)

	res, err := c.Post(uri, map[string]interface{}{
		"component_appid": c.GetComponentAppid(),
	})

	return res, err
}
