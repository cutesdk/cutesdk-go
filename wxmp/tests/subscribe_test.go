package tests

import (
	"fmt"
	"testing"
)

func TestSendSubscribeMsg(t *testing.T) {
	ins := getIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/message/subscribe/bizsend?access_token=%s", accessToken)

	params := map[string]interface{}{
		"touser":      "orNyi07T5GeE2nwVV7b1dJ3xGnPM",
		"template_id": "LTA5XpzSBWQQ3YI0hOw9Ucj-U1bvZIB_Ut9kHAhNBEk",
		"page":        "https://idoubi.cc",
		"data": map[string]interface{}{
			"thing5": map[string]interface{}{
				"value": "idoubi",
			},
			"thing7": map[string]interface{}{
				"value": "content",
			},
			"time2": map[string]interface{}{
				"value": "2021-12-02 13:11",
			},
		},
		"miniprogram": map[string]interface{}{
			"appid":    "wx25da2eca8fa3f4ef",
			"pagepath": "pages/index/index?scene=123",
		},
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetInt("errcode"), err)
}
