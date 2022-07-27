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
		Debug:   true,
	}

	ins, _ := wxwork.New(opts)

	ins.SetAccessToken("wdKvjEXLbqRo6iMgSsr3YlsIg7-FAMXegzMDKzOfjvDG9DYTYUhvWjXe_04MXR4iaQJggWWE8j-oFTvhv9ehAoxhogL1foJNkxsyXAerlI8Zx8AKw6ajHkZr0d-0Zsr_IbxeYLZJLN7SInQ2tbGSoCUo8JpTcjqghJtUsWgvWbuvB_IXqtwxMyCzcOvUvbI-c-8cLeJ-TpKfcV8O4rnwCg", 1*time.Hour)

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
		"name":     "你好朋友",
		"media_id": "3Vf19PeY1aL9sFAbyfzA9Mrz5t65R4XlRLbCjbr1yppJ5b_9IOMupCtYTM9YNBmdq",
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
		"open_kfid": "wkFtUHbgAAIgVvkwQgS88lYAFYl8mTZw",
		"name":      "大客服",
		"media_id":  "3Vf19PeY1aL9sFAbyfzA9Mrz5t65R4XlRLbCjbr1yppJ5b_9IOMupCtYTM9YNBmdq",
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
		"open_kfid": "wkFtUHbgAAu1dGdoGKf0Ggoos60IclKg",
		"scene":     "cae9846af84b52f56c1f74d8202c509b",
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
		"open_kfid": "wkFtUHbgAAaVpFUV8VL-0-pwn8OFAMWA",
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
		"open_kfid":          "wkFtUHbgAAaVpFUV8VL-0-pwn8OFAMWA",
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
