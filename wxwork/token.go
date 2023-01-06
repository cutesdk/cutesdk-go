package wxwork

import "github.com/cutesdk/cutesdk-go/common/request"

// FetchAccessToken: request get_access_token api
func (cli *Client) FetchAccessToken() (*request.Result, error) {
	uri := "/cgi-bin/gettoken"

	res, err := cli.Get(uri, map[string]interface{}{
		"corpid":     cli.opts.Corpid,
		"corpsecret": cli.opts.Secret,
	})

	return res, err
}
