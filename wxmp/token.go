package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// FetchAccessToken: request get_access_token api
func (ins *Instance) FetchAccessToken() (*request.Result, error) {
	uri := "/cgi-bin/token"

	res, err := ins.Get(uri, map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      ins.GetAppid(),
		"secret":     ins.GetSecret(),
	})

	return res, err
}

// FetchJsapiTicket: request get_jsapi_ticket api
func (ins *Instance) FetchJsapiTicket() (*request.Result, error) {
	accessToken, err := ins.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", accessToken)

	res, err := ins.Get(uri)

	return res, err
}

// FetchAuthorizerAccessToken: request get_authorizer_access_token api
func (ins *Instance) FetchAuthorizerAccessToken() (*request.Result, error) {
	provider := ins.GetAuthorizerProvider()
	if provider == nil {
		return nil, fmt.Errorf("no authorizer provider")
	}

	componentAccessToken, err := provider.GetComponentAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get component_access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/component/api_authorizer_token?component_access_token=%s", componentAccessToken)
	params := map[string]interface{}{
		"component_appid":          provider.GetComponentAppid(),
		"authorizer_appid":         ins.GetAppid(),
		"authorizer_refresh_token": ins.GetAuthorizerRefreshToken(),
	}
	res, err := provider.Post(uri, params)

	return res, err
}
