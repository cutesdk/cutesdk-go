package tests

import (
	"fmt"
	"testing"

	"github.com/cutesdk/cutesdk-go/wxwork"
)

func getContactIns() *wxwork.Instance {
	opts := &wxwork.Options{
		Corpid:  "xxx",
		Agentid: "contact",
		Secret:  "xxx",
		Debug:   true,
	}

	ins, _ := wxwork.New(opts)

	return ins
}

func TestGetDepartments(t *testing.T) {
	ins := getContactIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	pid := 0

	uri := fmt.Sprintf("/cgi-bin/department/list?access_token=%s&id=%d", accessToken, pid)

	res, err := ins.Get(uri)

	t.Error(res.GetArray("department"), err)
}

func TestGetUsers(t *testing.T) {
	ins := getContactIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	pid := 1
	fc := 1

	uri := fmt.Sprintf("/cgi-bin/user/simplelist?access_token=%s&department_id=%d&fetch_child=%d", accessToken, pid, fc)

	res, err := ins.Get(uri)

	t.Error(res.GetArray("userlist"), err)
}

func TestGetTags(t *testing.T) {
	ins := getContactIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("//cgi-bin/tag/list?access_token=%s", accessToken)

	res, err := ins.Get(uri)

	t.Error(res.GetArray("taglist"), err)
}
