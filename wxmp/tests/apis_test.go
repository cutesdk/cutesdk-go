package tests

import (
	"encoding/json"
	"testing"
)

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

func TestGetJssdkConfig(t *testing.T) {
	client := getClient()

	url := "https://idoubi.cc?p=123"

	res, err := client.GetJssdkConfig(url)

	if err != nil {
		t.Fatalf("get jssdk config failed: %v", err)
	}

	j, _ := json.Marshal(res)

	t.Errorf("%s", j)
}

func TestCreateQrcode(t *testing.T) {
	client := getClient()

	params := map[string]interface{}{
		"action_name": "QR_LIMIT_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": "qid=111",
			},
		},
	}

	res, err := client.CreateQrcode(params)

	t.Error(res.Get("long_url"), err)
}
