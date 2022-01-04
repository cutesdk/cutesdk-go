package wxopen

import (
	"testing"
)

func TestApiGet(t *testing.T) {
	sdk := getWxOpen()

	apiPath := "/wxa/gettemplatedraftlist"

	res, err := sdk.ApiGet(apiPath, map[string]string{
		"access_token": ComponentAccessToken,
	})

	if err != nil {
		t.Fatalf("request api error: %v\n", err)
	}

	if res.Parsed().Get("errcode").Int() != 0 {
		t.Fatalf("request api failed: %s", res)
	}
}

func TestApiPost(t *testing.T) {
	sdk := getWxOpen()

	apiPath := "/cgi-bin/component/get_domain_confirmfile"

	res, err := sdk.ApiPost(apiPath, map[string]string{
		"access_token": ComponentAccessToken,
	}, nil)

	if err != nil {
		t.Fatalf("request api error: %v\n", err)
	}

	if res.Parsed().Get("errcode").Int() != 0 {
		t.Fatalf("request api failed: %s", res)
	}
}
