package tests

import "testing"

func TestPostWithTokenInHeader(t *testing.T) {
	ins := getIns()

	uri := "/api/v2/tags/text/antidirt"

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	headers := map[string]interface{}{
		"X-Token": accessToken,
	}

	data := map[string]interface{}{
		"tasks": []map[string]string{
			{"content": "讲个笑话"},
		},
	}

	res, err := ins.Post(uri, data, headers)

	t.Error(res.GetString("log_id"), err)
}

func TestPostWithTokenInBody(t *testing.T) {
	ins := getIns()

	uri := "/api/apps/censor/image"

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	data := map[string]interface{}{
		"app_id":       ins.GetAppid(),
		"access_token": accessToken,
		"image_data":   "data:image/svg+xml;base64,xxx",
	}

	res, err := ins.Post(uri, data)

	t.Error(res.GetInt("error"), err)
}
