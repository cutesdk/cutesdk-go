package wxopen

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/cutesdk/cutesdk-go/wxmp"
)

// WxmpClient: wxmp client authorized to wxopen
type WxmpClient struct {
	*wxmp.Client
	wxopenCli *Client
}

// NewWxmpClient: new wxmp client
func (cli *Client) NewWxmpClient(appid, refreshToken string) (*WxmpClient, error) {
	baseCli, err := wxmp.NewClient(&wxmp.Options{
		Appid:   appid,
		Debug:   cli.opts.Debug,
		Timeout: cli.opts.Timeout,
		Cache:   cli.opts.Cache,
	})
	if err != nil {
		return nil, err
	}

	// set default authorizer_access_token cache key
	authorizerTokenCacheKey := fmt.Sprintf("wxopen.authorizer_access_token.%s.%s", cli.GetAppid(), appid)
	cli.SetAuthorizerTokenCacheKey(authorizerTokenCacheKey)

	// set default authorizer_access_token handler
	baseCli.SetAccessTokenHandler(NewAuthorizerToken(cli, appid, refreshToken))

	// set default jsapi_ticket cache key
	jsapiTicketCacheKey := fmt.Sprintf("wxopen.authorizer_jsapi_ticket.%s.%s", cli.GetAppid(), appid)
	baseCli.SetJsapiTicketCacheKey(jsapiTicketCacheKey)

	wxmpCli := &WxmpClient{baseCli, cli}

	return wxmpCli, nil
}

// GetOauthToken: get oauth access_token
func (cli *WxmpClient) GetOauthToken(code string) (*request.Result, error) {
	accessToken, err := cli.wxopenCli.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", token.ErrGetTokenFailed, err)
	}

	uri := "/sns/oauth2/component/access_token"
	params := map[string]interface{}{
		"appid":                  cli.GetAppid(),
		"code":                   code,
		"grant_type":             "authorization_code",
		"component_appid":        cli.wxopenCli.GetAppid(),
		"component_access_token": accessToken,
	}

	res, err := cli.Get(uri, params)

	return res, err
}
