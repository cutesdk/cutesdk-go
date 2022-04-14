package tests

import "testing"

func TestGetOauthUrl(t *testing.T) {
	client := getClient()

	redirectUri := "https://www.thinkwx.com/api/callback/"
	scope := "snsapi_userinfo"
	state := "test123"

	oauthUrl, err := client.GetOauthUrl(redirectUri, scope, state)

	t.Error(oauthUrl, err)
}

func TestGetOauthTok(t *testing.T) {
	client := getClient()

	code := "011kbs0006qsEN1NAn300GzQSs1kbs0F"

	res, err := client.GetOauthToken(code)

	t.Error(res, err)
}

func TestGetOauthUser(t *testing.T) {
	client := getClient()

	oauthAccessToken := "55_6JFl_DZXckzLFiQTGJruCNRllddreizk0cNHAZByZw9K7IDgbjQSKBVFGz_wehOUHt0Ayno9eVyiOrsudZG-rA"
	openid := "orNyi07T5GeE2nwVV7b1dJ3xGnPM"

	res, err := client.GetOauthUser(oauthAccessToken, openid)

	t.Error(res, err)
}
