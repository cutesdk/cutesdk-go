package tests

import "testing"

func TestFetchServiceAccessToken(t *testing.T) {
	client := getServiceClient()
	res, err := client.FetchAccessToken()

	if err != nil {
		t.Fatalf("fetch access_token error: %v\n", err)
	}

	t.Error(res, err)
}

func TestGetServiceAccessToken(t *testing.T) {
	client := getServiceClient()
	res, err := client.GetAccessToken()

	t.Error(res, err)
}
