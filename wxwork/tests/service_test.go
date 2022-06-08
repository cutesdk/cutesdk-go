package tests

import (
	"github.com/cutesdk/cutesdk-go/wxwork"
)

func getServiceIns() *wxwork.Instance {
	opts := &wxwork.Options{
		Corpid:  "wwa3f1494ad3d3713d",
		Agentid: "service",
		Secret:  "44d2imiTA4EhySj2TVfstu6LRh4dyGZef8oQcb43n_Y",
	}

	ins, _ := wxwork.New(opts)

	return ins
}

// func TestAddServiceAccount(t *testing.T) {
// 	client := getServiceIns()

// 	uri := fmt.Sprintf("/cgi-bin/kf/account/add?access_token=%s", accessToken)

// 	params := map[string]interface{}{
// 		"name":     "朋友坦白局",
// 		"media_id": "3-F4UoeyubrCGY2Jk3nruP5QTygxUZj-nlV8Ggrr4IQvZ49MtLrQnvaDie47ykJ5V",
// 	}

// 	res, err := client.Post(uri, params)

// 	t.Error(res, err)
// }

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

// func TestGetServiceAccounts(t *testing.T) {
// 	client := getServiceIns()

// 	uri := fmt.Sprintf("/cgi-bin/kf/account/list?access_token=%s", accessToken)

// 	params := map[string]interface{}{
// 		"offset": 0,
// 		"limit":  100,
// 	}

// 	res, err := client.Post(uri, params)

// 	t.Error(res, err)
// }

// func TestCreateServiceUrl(t *testing.T) {
// 	client := getServiceIns()

// 	uri := fmt.Sprintf("/cgi-bin/kf/add_contact_way?access_token=%s", accessToken)

// 	params := map[string]interface{}{
// 		"open_kfid": "wkFtUHbgAAvdl2TSXtUkCdvEtrYh66jQ",
// 		"scene":     "678fc75db2cc8f7f3e314c3399ed82f7",
// 	}

// 	res, err := client.Post(uri, params)

// 	if err != nil {
// 		t.Fatalf("request failed: %v", err)
// 	}

// 	if resUrl := res.GetString("url"); resUrl != "" {
// 		sceneParams := "/xzt/tbj/index?appId=490"
// 		serviceUrl := resUrl + "&scene_param=" + url.QueryEscape(sceneParams)
// 		t.Error("serviceUrl: ", serviceUrl)
// 	}

// 	t.Error(res, err)
// }

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
