package tests

import (
	"encoding/json"
	"testing"
)

func TestGetJssdkConfig(t *testing.T) {
	ins := getIns()

	url := "https://xxx.com?p=123"

	res, err := ins.GetJssdkConfig(url)

	if err != nil {
		t.Fatalf("get jssdk config failed: %v", err)
	}

	j, _ := json.Marshal(res)

	t.Errorf("%s", j)
}
