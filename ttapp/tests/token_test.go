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

	t.Error(res.GetString("data.access_token"), err)
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
