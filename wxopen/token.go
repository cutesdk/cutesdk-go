package wxopen

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// FetchAccessToken: request api to fetch component_access_token
func (cli *Client) FetchAccessToken() (*request.Result, error) {
	verifyTicket, err := cli.GetVerifyTicket()
	if err != nil {
		return nil, fmt.Errorf("get component_verify_ticket failed: %v", err)
	}

	uri := "/cgi-bin/component/api_component_token"

	res, err := cli.Post(uri, map[string]interface{}{
		"component_verify_ticket": verifyTicket,
		"component_appid":         cli.GetAppid(),
		"component_appsecret":     cli.GetSecret(),
	})

	return res, err
}

// FetchAuthorizerToken: request get_authorizer_access_token api
func (cli *Client) FetchAuthorizerToken(appid, refreshToken string) (*request.Result, error) {
	accessToken, err := cli.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("get component_access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/component/api_authorizer_token?component_access_token=%s", accessToken)
	params := map[string]interface{}{
		"component_appid":          cli.GetAppid(),
		"authorizer_appid":         appid,
		"authorizer_refresh_token": refreshToken,
	}

	res, err := cli.Post(uri, params)

	return res, err
}
