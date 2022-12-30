package ttapp

import "github.com/cutesdk/cutesdk-go/common/request"

// Login: ttapp user login
func (cli *Client) Login(code, anonymousCode string) (*request.Result, error) {
	uri := "/api/apps/v2/jscode2session"
	params := map[string]interface{}{
		"appid":  cli.GetAppid(),
		"secret": cli.GetSecret(),
		"code":   code,
	}
	if anonymousCode != "" {
		params["anonymous_code"] = anonymousCode
	}

	res, err := cli.Post(uri, params)

	return res, err
}
