package wxmp

import (
	"fmt"
	"net/url"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// GetOauthUrl: get oauth url
func (c *Client) GetOauthUrl(redirectUri, scope, state string) (*url.URL, error) {
	if redirectUri == "" {
		return nil, fmt.Errorf("invalid redirectUri")
	}
	if scope != "snsapi_base" && scope != "snsapi_userinfo" {
		return nil, fmt.Errorf("invalid scope")
	}

	oauthBaseUri := "https://open.weixin.qq.com/connect/oauth2/authorize"
	redirectUri = url.QueryEscape(redirectUri)

	oauthUrl := fmt.Sprintf("%s?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect", oauthBaseUri, c.GetAppid(), redirectUri, scope, state)

	return url.Parse(oauthUrl)
}

// GetOauthToken: get oauth access_token
func (c *Client) GetOauthToken(code string) (request.Result, error) {
	uri := fmt.Sprintf("/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", c.GetAppid(), c.GetSecret(), code)

	res, err := c.Get(uri)

	return res, err
}

// GetOauthUser: get oauth userinfo
func (c *Client) GetOauthUser(oauthAccessToken, openid string) (request.Result, error) {
	uri := fmt.Sprintf("/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", oauthAccessToken, openid)

	res, err := c.Get(uri)

	return res, err
}
