package ttapp

import "github.com/cutesdk/cutesdk-go/common/request"

// FetchAccessToken: request get_access_token api
func (cli *Client) FetchAccessToken() (*request.Result, error) {
	apiPath := "/api/apps/v2/token"

	res, err := cli.Post(apiPath, map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      cli.GetAppid(),
		"secret":     cli.GetSecret(),
	})

	return res, err
}
