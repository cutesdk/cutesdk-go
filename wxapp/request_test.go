package wxapp

import (
	"testing"
)

func TestApiGet(t *testing.T) {
	sdk := getSdk()

	apiPath := "/wxa/getnearbypoilist"

	res, err := sdk.ApiGet(apiPath, map[string]string{
		"page":         "1",
		"page_rows":    "10",
		"access_token": AccessToken,
	})

	if err != nil {
		t.Fatalf("request api error: %v\n", err)
	}

	if res.Parsed().Get("errcode").Int() != 0 {
		t.Fatalf("request api failed: %s", res)
	}
}

func TestApiPost(t *testing.T) {
	sdk := getSdk()

	apiPath := "/datacube/getweanalysisappiddailyretaininfo"

	res, err := sdk.ApiPost(apiPath, map[string]string{
		"access_token": AccessToken,
	}, map[string]interface{}{
		"begin_date": "20220101",
		"end_date":   "20220101",
	})

	if err != nil {
		t.Fatalf("request api error: %v\n", err)
	}

	if res.Parsed().Get("errcode").Int() != 0 {
		t.Fatalf("request api failed: %s", res)
	}

	t.Error(res)
}

func getSdk() *WxApp {
	opts := Options{
		Debug:     true,
		Appid:     "wx25da2eca8fa3f4ef",
		AppSecret: "1324f564b26f9f8006515e13660876ef",
	}

	sdk := New(opts)

	return sdk
}
