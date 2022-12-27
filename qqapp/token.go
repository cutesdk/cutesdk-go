package qqapp

import "github.com/cutesdk/cutesdk-go/common/request"

// FetchAccessToken: request get_access_token api
func (cli *Client) FetchAccessToken() (*request.Result, error) {
	uri := "/api/getToken"

	res, err := cli.Get(uri, map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      cli.GetAppid(),
		"secret":     cli.GetSecret(),
	})

	return res, err
}
