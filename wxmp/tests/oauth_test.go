package tests

import "testing"

func TestGetOauthUrl(t *testing.T) {
	ins := getIns()

	redirectUri := "https://www.xxx.com/api/callback/"
	scope := "snsapi_userinfo"
	state := "test123"

	oauthUrl, err := ins.GetOauthUrl(redirectUri, scope, state)

	t.Error(oauthUrl, err)
}

func TestGetOauthToken(t *testing.T) {
	ins := getIns()

	code := "xxx"

	res, err := ins.GetOauthToken(code)

	t.Error(res.GetInt("errcode"), err)
}

func TestGetOauthUser(t *testing.T) {
	ins := getIns()

	oauthAccessToken := "xxx-xxx"
	openid := "xxx"

	res, err := ins.GetOauthUser(oauthAccessToken, openid)

	t.Error(res.GetInt("errcode"), err)
}
