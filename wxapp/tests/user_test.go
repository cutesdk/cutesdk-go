package tests

import (
	"errors"
	"testing"

	"github.com/cutesdk/cutesdk-go/common/token"
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

	uri := "/wxa/business/getuserphonenumber"

	params := map[string]interface{}{
		"code": "xxx",
	}

	res, err := ins.PostWithAccessToken(uri, params)

	if err != nil {
		if errors.Is(err, token.ErrGetTokenFailed) {
			t.Fatalf("%v", err)
		}
	}
	t.Error(res.GetInt("errcode"), err)
}
