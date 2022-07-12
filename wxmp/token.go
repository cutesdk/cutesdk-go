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
