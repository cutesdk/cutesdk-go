package tests

import (
	"fmt"
	"testing"
)

func TestGetUserList(t *testing.T) {
	ins := getIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/user/get?access_token=%s", accessToken)

	params := map[string]interface{}{
		"next_openid": "",
	}

	res, err := ins.Get(uri, params)

	t.Error(res.GetInt("total"), err)
}

func TestGetUserInfo(t *testing.T) {
	ins := getIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/user/info?access_token=%s", accessToken)

	params := map[string]interface{}{
		"openid": "xxx",
		"lang":   "zh_CN",
	}

	res, err := ins.Get(uri, params)

	t.Error(res.GetString("openid"), err)
}
