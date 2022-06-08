package tests

import "testing"

func TestGetOauthUrl(t *testing.T) {
	ins := getIns()

	redirectUri := "https://www.thinkwx.com/api/callback/"
	scope := "snsapi_userinfo"
	state := "test123"

	oauthUrl, err := ins.GetOauthUrl(redirectUri, scope, state)

	t.Error(oauthUrl, err)
}

func TestGetOauthToken(t *testing.T) {
	ins := getIns()

	code := "011kbs0006qsEN1NAn300GzQSs1kbs0F"

	res, err := ins.GetOauthToken(code)

	t.Error(res.GetInt("errcode"), err)
}

func TestGetOauthUser(t *testing.T) {
	ins := getIns()

	oauthAccessToken := "55_6JFl_DZXckzLFiQTGJruCNRllddreizk0cNHAZByZw9K7IDgbjQSKBVFGz_wehOUHt0Ayno9eVyiOrsudZG-rA"
	openid := "orNyi07T5GeE2nwVV7b1dJ3xGnPM"

	res, err := ins.GetOauthUser(oauthAccessToken, openid)

	t.Error(res.GetInt("errcode"), err)
}
