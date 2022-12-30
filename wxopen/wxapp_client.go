package wxopen

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/common/token"
	"github.com/cutesdk/cutesdk-go/wxapp"
)

// WxappClient: wxapp client authorized to wxopen
type WxappClient struct {
	*wxapp.Client
	wxopenCli *Client
}

// NewWxappClient: new wxapp client
func (cli *Client) NewWxappClient(appid, refreshToken string) (*WxappClient, error) {
	baseCli, err := wxapp.NewClient(&wxapp.Options{
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

	wxappCli := &WxappClient{baseCli, cli}

	return wxappCli, nil
}

// Login: wxapp user login
func (wxappCli *WxappClient) Login(code string) (*request.Result, error) {
	accessToken, err := wxappCli.wxopenCli.GetAccessToken()
	if err != nil {
		return nil, fmt.Errorf("%w: %v", token.ErrGetTokenFailed, err)
	}

	uri := "/sns/component/jscode2session"
	params := map[string]interface{}{
		"component_access_token": accessToken,
		"appid":                  wxappCli.GetAppid(),
		"grant_type":             "authorization_code",
		"component_appid":        wxappCli.wxopenCli.GetAppid(),
		"js_code":                code,
	}

	res, err := wxappCli.Get(uri, params)

	return res, err
}
