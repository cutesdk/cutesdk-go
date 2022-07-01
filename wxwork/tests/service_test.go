package tests

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/cutesdk/cutesdk-go/wxwork"
)

func getServiceIns() *wxwork.Instance {
	opts := &wxwork.Options{
		Corpid:  "xxx",
		Agentid: "service",
		Secret:  "xxx",
	}

	ins, _ := wxwork.New(opts)

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
		"name":     "小客服",
		"media_id": "xxx-xxx",
	}

	res, err := ins.Post(uri, params)

	t.Error(res, err)
}

func TestUpdateServiceAccount(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/account/update?access_token=%s", accessToken)

	params := map[string]interface{}{
		"open_kfid": "xxx",
		"name":      "大客服",
		"media_id":  "xxx-xxx",
	}

	res, err := ins.Post(uri, params)

	t.Error(res, err)
}

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
		"open_kfid": "xxx",
		"scene":     "oid=3",
	}

	res, err := ins.Post(uri, params)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resUrl := res.GetString("url"); resUrl != "" {
		sceneParams := "/index?appId=123"
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
		"open_kfid": "xxx",
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
		"open_kfid":          "xxx",
		"department_id_list": []int64{1},
	}

	res, err := ins.Post(uri, params)

	t.Error(res, err)
}

func TestSyncMsg(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/sync_msg?access_token=%s", accessToken)

	params := map[string]interface{}{
		"cursor": "",
		"token":  "",
	}

	res, err := ins.Post(uri, params)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	t.Error(res, err)
}

func TestSendMsg(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/send_msg?access_token=%s", accessToken)

	params := map[string]interface{}{
		"touser":    "xxx",
		"open_kfid": "xxx",
		"msgtype":   "text",
		"text": map[string]interface{}{
			"content": "你好啊",
		},
	}

	res, err := ins.Post(uri, params)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	t.Error(res, err)
}

func TestBatchGetCustomer(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/customer/batchget?access_token=%s", accessToken)

	params := map[string]interface{}{
		"external_userid_list":       []string{"xxx"},
		"need_enter_session_context": 1,
	}

	res, err := ins.Post(uri, params)

	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	t.Error(res, err)
}

func TestServiceState(t *testing.T) {
	ins := getServiceIns()

	accessToken, err := ins.GetAccessToken()
	if err != nil {
		t.Fatalf("get access_token failed: %v", err)
	}

	uri := fmt.Sprintf("/cgi-bin/kf/get_corp_qualification?access_token=%s", accessToken)

	res, err := ins.Get(uri)

	t.Error(res, err)
}
