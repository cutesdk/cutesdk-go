package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/cutesdk/cutesdk-go/wxwork"
)

func getContactIns() *wxwork.Instance {
	opts := &wxwork.Options{
		Corpid:  "wwa3f1494ad3d3713d",
		Agentid: "contact",
		Secret:  "hz0ydrPX5ZjG_yDLyRk1xzt-tALpgYE9-sMs8Uk7MqE",
		Debug:   true,
	}

	ins, _ := wxwork.New(opts)

	accessToken := `QBHr-JGerArt6tfG3hYNSpMCTS2aCFcZCXQjBPkPAPXAqkcW_tTJMKmYhaeMb3usaPEfLe1pUT5NHfIMh0gI6O6GcPlvK_xyCNYZNM2AHKH76jwdcjxoWrKSWO7gawOPkzsXUBmki0oVB5PjlKmDHoz-nNzYWbkkTpTZsD5Ii01FdDY8TdyorfE2a80kv5SiRu8-N8Z4scaN7PKD-CxE3Q`
	ins.SetAccessToken(accessToken, 5*time.Second)

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
