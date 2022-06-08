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
		"js_code":    "051kk0nl2jfTd94EqBnl2cp1IS3kk0nO",
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
		"code": "cda802f413a0a7c9767819742a84f3fe3e7839da6661e62f39d0290a21116dce",
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetInt("errcode"), err)
}
