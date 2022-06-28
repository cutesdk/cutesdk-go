package tests

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/cutesdk/cutesdk-go/wxwork"
)

func getServiceIns() *wxwork.Instance {
	opts := &wxwork.Options{
		Corpid:  "ww8b7343f7397a7d22",
		Agentid: "service",
		Secret:  "q3PwXW-HK3Bo3uqVyvQT6dMvmtxfLrP04j_bXvFUQJc",
	}

	ins, _ := wxwork.New(opts)

	accessToken := `oJxtjOStctm3Pguy8LeOY6wo606SLLwAGuCQKdaVhNcyeMbvFUl-VGnjH4zou_H2yQj-P2DMXD_0q2GegKkZ1fGxtDI_VQLCmkM_FcHCv6Bz4tLCKWpv0Wnn-L63yYbA5eK4SpYSNHR56-dokCAN-GLsUGjOJSdqUa160QxZKUYmBGFBri02IE9LqogWHbSJehL1gpkVr-i05Q3xihcjtQ`
	ins.SetAccessToken(accessToken, 5*time.Second)

	return ins
}

func TestAddServiceAccount(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/account/add?access_token=%s", accessToken)

	params := map[string]interface{}{
		"name":     "星座配对",
		"media_id": "3tyzTX66GVgYO_EuodDAdtRi3OhbdDynFvpHQoohfllwcM59PEavO-jFfpjgV81VI",
	}

	res, err := ins.Post(uri, params)

	t.Error(res, err)
}

// func TestUpdateServiceAccount(t *testing.T) {
// 	client := getServiceIns()

// 	uri := fmt.Sprintf("/cgi-bin/kf/account/update?access_token=%s", accessToken)

// 	params := map[string]interface{}{
// 		"open_kfid": "wkFtUHbgAAK8PQFErEjUJXMHV_yKxKFg",
// 		"name":      "大冒险",
// 		"media_id":  "38FfWtorugNe3skroEG8bRqvl_gLlJUx5b1sdVtbGbVVijxCpEWQV_9sBVASbIRaJ",
// 	}

// 	res, err := client.Post(uri, params)

// 	t.Error(res, err)
// }

func TestGetServiceAccounts(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/account/list?access_token=%s", accessToken)

	params := map[string]interface{}{
		"offset": 0,
		"limit":  100,
	}

	res, err := ins.Post(uri, params)

	t.Error(res, err)
}

func TestCreateServiceUrl(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/add_contact_way?access_token=%s", accessToken)

	params := map[string]interface{}{
		"open_kfid": "wkFtUHbgAAM5Znc3_gNsSb8z7gzUhVew",
		"scene":     "921fb0e2f3b4864fc7fb885292da7170",
	}

	res, err := ins.Post(uri, params)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resUrl := res.GetString("url"); resUrl != "" {
		sceneParams := "/index?appId=22001"
		serviceUrl := resUrl + "&scene_param=" + url.QueryEscape(sceneParams)
		t.Error("serviceUrl: ", serviceUrl)
	}

	t.Error(res, err)
}

func TestGetServicers(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/servicer/list?access_token=%s", accessToken)

	params := map[string]interface{}{
		"open_kfid": "wkRdKcDgAAm_OIuPQCvuANOX8p_xlG_A",
	}

	res, err := ins.Get(uri, params)

	t.Error(res, err)
}

func TestAddServicer(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/servicer/add?access_token=%s", accessToken)

	params := map[string]interface{}{
		"open_kfid":          "wkRdKcDgAAm_OIuPQCvuANOX8p_xlG_A",
		"department_id_list": []int64{1},
	}

	res, err := ins.Post(uri, params)

	t.Error(res, err)
}

// func TestSyncMsg(t *testing.T) {
// 	client := getServiceIns()

// 	uri := fmt.Sprintf("/cgi-bin/kf/sync_msg?access_token=%s", accessToken)

// 	params := map[string]interface{}{
// 		"cursor": "",
// 		"token":  "",
// 	}

// 	res, err := client.Post(uri, params)

// 	if err != nil {
// 		t.Fatalf("request failed: %v", err)
// 	}

// 	t.Error(res, err)
// }

// func TestSendMsg(t *testing.T) {
// 	client := getServiceIns()

// 	uri := fmt.Sprintf("/cgi-bin/kf/send_msg?access_token=%s", accessToken)

// 	params := map[string]interface{}{
// 		"touser":    "wmRdKcDgAAYRYwpe9Il3oHBAHWilIqFg",
// 		"open_kfid": "wkRdKcDgAARzHX7ezeLTM74ILESuKIiQ",
// 		"msgtype":   "text",
// 		"text": map[string]interface{}{
// 			"content": "你好啊",
// 		},
// 	}

// 	res, err := client.Post(uri, params)

// 	if err != nil {
// 		t.Fatalf("request failed: %v", err)
// 	}

// 	t.Error(res, err)
// }

// func TestBatchGetCustomer(t *testing.T) {
// 	client := getServiceIns()

// 	uri := fmt.Sprintf("/cgi-bin/kf/customer/batchget?access_token=%s", accessToken)

// 	params := map[string]interface{}{
// 		"external_userid_list":       []string{"wmRdKcDgAAYRYwpe9Il3oHBAHWilIqFg"},
// 		"need_enter_session_context": 1,
// 	}

// 	res, err := client.Post(uri, params)

// 	if err != nil {
// 		t.Fatalf("request failed: %v", err)
// 	}

// 	t.Error(res, err)
// }

// func TestServiceState(t *testing.T) {
// 	client := getServiceIns()

// 	uri := fmt.Sprintf("/cgi-bin/kf/get_corp_qualification?access_token=%s", accessToken)

// 	res, err := client.Get(uri)

// 	t.Error(res, err)
// }
