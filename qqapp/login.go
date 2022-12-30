package qqapp

import "github.com/cutesdk/cutesdk-go/common/request"

// Login: qqapp user login
func (cli *Client) Login(code string) (*request.Result, error) {
	uri := "/sns/jscode2session"
	params := map[string]interface{}{
		"appid":      cli.GetAppid(),
		"secret":     cli.GetSecret(),
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	res, err := cli.Get(uri, params)

	return res, err
}
