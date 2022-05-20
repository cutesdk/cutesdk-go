package tests

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGetServiceAccounts(t *testing.T) {
	client := getServiceClient()

	accessToken, err := client.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/account/list?access_token=%s", accessToken)

	params := map[string]interface{}{
		"offset": 0,
		"limit":  100,
	}

	res, err := client.Post(uri, params)

	t.Error(res, err)
}

func TestCreateServiceUrl(t *testing.T) {
	client := getServiceClient()

	accessToken, err := client.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/add_contact_way?access_token=%s", accessToken)

	params := map[string]interface{}{
		"open_kfid": "wkRdKcDgAARzHX7ezeLTM74ILESuKIiQ",
		"scene":     "b90f32032110e60fd74e751f8a016505",
	}

	res, err := client.Post(uri, params)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resUrl := res.GetString("url"); resUrl != "" {
		sceneParams := "qid=363&openid=test123"
		serviceUrl := resUrl + "&scene_param=" + url.QueryEscape(sceneParams)
		t.Error("serviceUrl: ", serviceUrl)
	}

	t.Error(res, err)
}
