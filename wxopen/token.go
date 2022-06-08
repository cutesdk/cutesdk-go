package wxopen

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// FetchComponentAccessToken: request api to fetch component_access_token
func (ins *Instance) FetchComponentAccessToken() (*request.Result, error) {
	componentVerifyTicket, err := ins.GetComponentVerifyTicket()
	if err != nil {
		return nil, fmt.Errorf("get component_verify_ticket failed: %v", err)
	}

	uri := "/cgi-bin/component/api_component_token"

	res, err := ins.Post(uri, map[string]interface{}{
		"component_verify_ticket": componentVerifyTicket,
		"component_appid":         ins.GetComponentAppid(),
		"component_appsecret":     ins.GetComponentAppsecret(),
	})

	return res, err
}
