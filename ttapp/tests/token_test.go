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
