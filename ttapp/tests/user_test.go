package tests

import "testing"

func TestCode2Session(t *testing.T) {
	ins := getIns()

	uri := "/api/apps/v2/jscode2session"

	code := "xxx"
	anonymousCode := "xxx"

	params := map[string]interface{}{
		"appid":          ins.GetAppid(),
		"secret":         ins.GetSecret(),
		"code":           code,
		"anonymous_code": anonymousCode,
	}

	res, err := ins.Post(uri, params)

	t.Error(res.GetMap("data"), err)
}
