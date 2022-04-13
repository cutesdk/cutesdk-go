package tests

import (
	"fmt"
	"testing"
)

func TestApiGet(t *testing.T) {
	client := getClient()

	accessToken, err := client.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", accessToken)
	}

	uri := fmt.Sprintf("/wxa/getnearbypoilist?access_token=%s", accessToken)

	res, err := client.Get(uri, map[string]interface{}{
		"page":      "1",
		"page_rows": "10",
	})

	t.Error(res, err)
}

func TestApiPost(t *testing.T) {
	client := getClient()

	accessToken, err := client.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", accessToken)
	}

	uri := fmt.Sprintf("/datacube/getweanalysisappiddailyretaininfo?access_token=%s", accessToken)

	res, err := client.Post(uri, map[string]interface{}{
		"begin_date": "20220101",
		"end_date":   "20220101",
	})

	t.Error(res, err)
}
