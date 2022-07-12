package tests

import (
	"testing"
)

func TestFetchAccessToken(t *testing.T) {
	ins := getIns()
	res, err := ins.FetchAccessToken()

	if err != nil {
		t.Fatalf("fetch access_token error: %v\n", err)
	}

	t.Error(res.GetString("access_token"), err)
}

func TestGetAccessToken(t *testing.T) {
	ins := getIns()
	res, err := ins.GetAccessToken()

	t.Error(res, err)
}

func TestRefreshAccessToken(t *testing.T) {
	ins := getIns()
	res, err := ins.RefreshAccessToken()

	t.Error(res, err)
}

func TestFetchJsapiTicket(t *testing.T) {
	ins := getIns()

	res, err := ins.FetchJsapiTicket()

	t.Error(res.GetString("ticket"), err)
}

func TestGetJsapiTicket(t *testing.T) {
	ins := getIns()

	res, err := ins.GetJsapiTicket()

	t.Error(res, err)
}

func TestRefreshJsapiTicket(t *testing.T) {
	ins := getIns()

	res, err := ins.RefreshJsapiTicket()

	t.Error(res, err)
}
