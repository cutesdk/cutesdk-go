package wxopen

import (
	"fmt"
	"net/url"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// GetOauthUrl: get oauth url
func (cli *Client) GetOauthUrl(redirectUri, scope string, extra map[string]string) (string, error) {
	oauthUrl := fmt.Sprintf("https://open.weixin.qq.com/connect/qrconnect?appid=%s&redirect_uri=%s&response_type=code&scope=%s",
		cli.GetAppid(),
		url.QueryEscape(redirectUri),
		scope,
	)
	for k, v := range extra {
		oauthUrl += fmt.Sprintf("&%s=%s", k, v)
	}

	oauthUrl += "#wechat_redirect"

	return oauthUrl, nil
}

// GetOauthToken: get oauth access_token
func (cli *Client) GetOauthToken(code string) (*request.Result, error) {
	uri := fmt.Sprintf("/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		cli.GetAppid(),
		cli.GetSecret(),
		code,
	)

	res, err := cli.Get(uri)

	return res, err
}

// RefreshOauthToken: refresh oauth access_token
func (cli *Client) RefreshOauthToken(refreshToken string) (*request.Result, error) {
	uri := fmt.Sprintf("/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
		cli.GetAppid(),
		refreshToken,
	)

	res, err := cli.Get(uri)

	return res, err
}

// GetOauthUser: get oauth userinfo
func (cli *Client) GetOauthUser(oauthAccessToken, openid string) (*request.Result, error) {
	uri := fmt.Sprintf("/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", oauthAccessToken, openid)

	res, err := cli.Get(uri)

	return res, err
}
