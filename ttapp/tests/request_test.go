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
		"image_data":   "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjgiIGhlaWdodD0iMjgiIHZpZXdCb3g9IjAgMCAyOCAyOCIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTE1LjgwMzIgMTkuOTk5NUgxMS45NzY2VjI2LjAwMDJIMTUuODAzMlYxOS45OTk1WiIgZmlsbD0iIzAwQzhEMiIvPgo8cGF0aCBkPSJNMTUuODAyOSAySDExLjk3NjNWMTIuOTUwNkwyLjAwMjkzIDE4LjQyNThMMy45MTYyNCAyMS41NzI4TDEzLjg4OTYgMTYuMTAwMkwyMy44NjU4IDIxLjU3MjhMMjUuNzc2MyAxOC40MjU4TDE1LjgwMjkgMTIuOTUwNlYyWiIgZmlsbD0iIzNDODlGRiIvPgo8cGF0aCBkPSJNMy45MTMzMiA2LjQxOTk0TDIgOS41NzAzMUw3LjQ2NjU3IDEyLjU3MDZMOS4zNzk4OSA5LjQyMDI3TDMuOTEzMzIgNi40MTk5NFoiIGZpbGw9IiMwMEM4RDIiLz4KPHBhdGggZD0iTTIzLjg2NSA2LjQyMTU0TDE4LjM5ODQgOS40MjE4OEwyMC4zMTE4IDEyLjU3MjNMMjUuNzc4MyA5LjU3MTkyTDIzLjg2NSA2LjQyMTU0WiIgZmlsbD0iIzAwQzhEMiIvPgo8L3N2Zz4K",
	}

	res, err := ins.Post(uri, data)

	t.Error(res.GetInt("error"), err)
}
