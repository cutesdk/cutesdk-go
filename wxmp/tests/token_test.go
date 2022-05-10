package tests

import (
	"testing"
)

func TestFetchAccessToken(t *testing.T) {
	client := getClient()
	res, err := client.FetchAccessToken()

	if err != nil {
		t.Fatalf("fetch access_token error: %v\n", err)
	}

	t.Error(res, err)
}

func TestGetAccessToken(t *testing.T) {
	client := getClient()
	res, err := client.GetAccessToken()

	t.Error(res, err)
}

func TestRefreshAccessToken(t *testing.T) {
	client := getClient()
	res, err := client.RefreshAccessToken()

	t.Error(res, err)
}

func TestFetchJsapiTicket(t *testing.T) {
	client := getClient()

	res, err := client.FetchJsapiTicket()

	t.Error(res, err)
}

func TestGetJsapiTicket(t *testing.T) {
	client := getClient()

	res, err := client.GetJsapiTicket()

	t.Error(res, err)
}

func TestRefreshJsapiTicket(t *testing.T) {
	client := getClient()

	res, err := client.RefreshJsapiTicket()

	t.Error(res, err)
}
