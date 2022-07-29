package wxmp

import (
	"fmt"
	"net/url"

	"github.com/cutesdk/cutesdk-go/common/request"
)

// GetOauthUrl: get oauth url
func (ins *Instance) GetOauthUrl(redirectUri, scope, state string, args ...string) (*url.URL, error) {
	if redirectUri == "" {
		return nil, fmt.Errorf("invalid redirectUri")
	}
	if scope != "snsapi_base" && scope != "snsapi_userinfo" {
		return nil, fmt.Errorf("invalid scope")
	}

	oauthBaseUri := "https://open.weixin.qq.com/connect/oauth2/authorize"
	redirectUri = url.QueryEscape(redirectUri)

	params := url.Values{}
	params.Set("appid", ins.GetAppid())
	params.Set("redirect_uri", redirectUri)
	params.Set("response_type", "code")
	params.Set("scope", scope)
	params.Set("state", state)

	if len(args) > 0 {
		params.Set("forcePopup", args[0])
	}
	if len(args) > 1 {
		params.Set("forceSnapShot", args[1])
	}

	oauthUrl := fmt.Sprintf("%s?%s#wechat_redirect", oauthBaseUri, params.Encode())

	return url.Parse(oauthUrl)
}

// GetOauthToken: get oauth access_token
func (ins *Instance) GetOauthToken(code string) (*request.Result, error) {
	uri := fmt.Sprintf("/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", ins.GetAppid(), ins.GetSecret(), code)

	res, err := ins.Get(uri)

	return res, err
}

// GetOauthUser: get oauth userinfo
func (ins *Instance) GetOauthUser(oauthAccessToken, openid string) (*request.Result, error) {
	uri := fmt.Sprintf("/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN", oauthAccessToken, openid)

	res, err := ins.Get(uri)

	return res, err
}
