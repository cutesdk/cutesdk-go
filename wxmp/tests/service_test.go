package tests

import (
	"fmt"
	"testing"
)

func TestSendServiceMsg(t *testing.T) {
	ins := getIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/message/custom/send?access_token=%s", accessToken)

	data := map[string]interface{}{
		"touser":  "orNyi07T5GeE2nwVV7b1dJ3xGnPM",
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": "hello world",
		},
	}

	res, err := ins.Post(uri, data)

	t.Error(res.GetInt("errcode"), err)
}
