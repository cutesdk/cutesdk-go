package examples

import (
	"fmt"
	"log"
	"net/url"
)

func ExampleGetServiceAccounts() {
	cli := getServiceClient()

	uri := "/cgi-bin/kf/account/list"
	params := map[string]interface{}{
		"offset": 0,
		"limit":  100,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("account_list"))
	// Output: xxx
}

func ExampleAddServiceAccount() {
	cli := getServiceClient()

	uri := "/cgi-bin/kf/account/add"
	params := map[string]interface{}{
		"name":     "helper",
		"media_id": "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("open_kfid"))
	// Output: xxx
}

func ExampleUpdateServiceAccount() {
	cli := getServiceClient()

	uri := "/cgi-bin/kf/account/update"
	params := map[string]interface{}{
		"open_kfid": "xxx",
		"name":      "xxx",
		"media_id":  "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}

func ExampleCreateServiceUrl() {
	cli := getServiceClient()

	uri := "/cgi-bin/kf/add_contact_way"

	params := map[string]interface{}{
		"open_kfid": "xxx",
		"scene":     "123",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if resUrl := res.GetString("url"); resUrl != "" {
		sceneParams := "/index?appId=123"
		serviceUrl := resUrl + "&scene_param=" + url.QueryEscape(sceneParams)
		fmt.Println(serviceUrl)
	}

	// Output: xxx
}

func ExampleSyncServiceMsg() {
	cli := getServiceClient()

	uri := "/cgi-bin/kf/sync_msg"
	params := map[string]interface{}{
		"cursor":    "",
		"token":     "xxx",
		"limit":     1000,
		"open_kfid": "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("msg_list"))
	// Output: xxx
}

func ExampleSendServiceMsg() {
	cli := getServiceClient()

	uri := "/cgi-bin/kf/send_msg"
	params := map[string]interface{}{
		"touser":    "xxx",
		"open_kfid": "xxx",
		"msgtype":   "text",
		"text": map[string]interface{}{
			"content": "hello, my friend",
		},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}

func ExampleSendServiceWelcomeMsg() {
	cli := getServiceClient()

	uri := "/cgi-bin/kf/send_msg_on_event"
	params := map[string]interface{}{
		"code":    "xxx",
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": "hello, my friend",
		},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}

func ExampleGetServiceCustomers() {
	cli := getServiceClient()

	uri := "/cgi-bin/kf/customer/batchget"
	params := map[string]interface{}{
		"external_userid_list": []string{
			"xxx",
		},
		"need_enter_session_context": 1,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("customer_list"))
	// Output: xxx
}
