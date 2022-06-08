package tests

import (
	"fmt"
	"testing"
)

func TestGetCurrentMen(t *testing.T) {
	ins := getIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/get_current_selfmenu_info?access_token=%s", accessToken)

	res, err := ins.Get(uri)

	t.Error(res.GetArray("selfmenu_info.button"), err)
}
