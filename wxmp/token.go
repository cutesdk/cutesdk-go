package wxmp

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
)

// FetchAccessToken: request get_access_token api
func (cli *Client) FetchAccessToken() (*request.Result, error) {
	uri := "/cgi-bin/token"

	res, err := cli.Get(uri, map[string]interface{}{
		"grant_type": "client_credential",
		"appid":      cli.GetAppid(),
		"secret":     cli.GetSecret(),
	})

	return res, err
}

// FetchJsapiTicket: request get_jsapi_ticket api
func (cli *Client) FetchJsapiTicket() (*request.Result, error) {
	accessToken, err := cli.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", token.ErrGetTokenFailed, err)
	}

	uri := fmt.Sprintf("/cgi-bin/ticket/getticket?access_token=%s&type=jsapi", accessToken)

	res, err := cli.Get(uri)

	return res, err
}
