package tests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/cutesdk/cutesdk-go/common/token"
)

func TestGetUserList(t *testing.T) {
	ins := getIns()

	uri := "/cgi-bin/user/get"
	params := map[string]interface{}{
		"next_openid": "",
	}
	res, err := ins.GetWithToken(uri, params)

	if err != nil {
		if errors.Is(err, token.ErrGetTokenFailed) {
			t.Fatalf("%v", err)
		}
	}

	t.Error(res, err)
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
