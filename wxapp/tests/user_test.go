package tests

import (
	"fmt"
	"testing"
)

func TestCode2Session(t *testing.T) {
	ins := getIns()

	uri := "/sns/jscode2session"

	params := map[string]interface{}{
		"appid":      ins.GetAppid(),
		"secret":     ins.GetSecret(),
		"js_code":    "xxx",
		"grant_type": "authorization_code",
	}

	res, err := ins.Get(uri, params)

	t.Error(res.GetInt("errcode"), err)
}

func TestGetUserPhone(t *testing.T) {
	ins := getIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/wxa/business/getuserphonenumber?access_token=%s", accessToken)

	params := map[string]interface{}{
		"code": "xxx",
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetInt("errcode"), err)
}
