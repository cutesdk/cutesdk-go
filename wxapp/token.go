package wxapp

import "github.com/cutesdk/cutesdk-go/common/request"

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
