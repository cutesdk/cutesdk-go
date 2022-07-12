package ttapp

import "github.com/cutesdk/cutesdk-go/common/request"

// FetchAccessToken: request get_access_token api
func (ins *Instance) FetchAccessToken() (*request.Result, error) {
	apiPath := "/api/apps/v2/token"

	res, err := ins.Post(apiPath, map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      ins.GetAppid(),
		"secret":     ins.GetSecret(),
	})

	return res, err
}
