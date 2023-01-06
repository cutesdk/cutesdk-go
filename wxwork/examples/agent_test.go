package examples

import (
	"fmt"
	"log"
)

func ExampleGetAgentInfo() {
	cli := getAgentClient()

	uri := "/cgi-bin/agent/get"
	params := map[string]interface{}{
		"agentid": cli.GetAppid(),
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("name"))
	// Output: xxx
}

func ExampleGetAgentMenu() {
	cli := getAgentClient()

	uri := "/cgi-bin/menu/get"
	params := map[string]interface{}{
		"agentid": cli.GetAppid(),
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}

func ExampleCreateAgentMenu() {
	cli := getAgentClient()

	uri := fmt.Sprintf("/cgi-bin/menu/create?agentid=%s", cli.GetAppid())
	params := map[string]interface{}{
		"button": []map[string]interface{}{
			{
				"type": "click",
				"name": "今日歌曲",
				"key":  "V1001_TODAY_MUSIC",
			},
			{
				"name": "菜单",
				"sub_button": []map[string]interface{}{
					{
						"type": "view",
						"name": "搜索",
						"url":  "http://www.soso.com/",
					},
					{
						"type": "click",
						"name": "赞一下我们",
						"key":  "V1001_GOOD",
					},
				},
			},
		},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}

func ExampleDeleteAgentMenu() {
	cli := getAgentClient()

	uri := "/cgi-bin/menu/delete"
	params := map[string]interface{}{
		"agentid": cli.GetAppid(),
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}

func ExampleGetDepartments() {
	cli := getAgentClient()

	uri := "/cgi-bin/department/list"
	params := map[string]interface{}{
		"id": 0,
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("department"))
	// Output: xxx
}

func ExampleGetUsers() {
	cli := getAgentClient()

	uri := "/cgi-bin/user/list"
	params := map[string]interface{}{
		"department_id": 1,
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("userlist"))
	// Output: xxx
}

func ExampleSendAgentMsg() {
	cli := getAgentClient()

	uri := "/cgi-bin/message/send"
	params := map[string]interface{}{
		"agentid": cli.GetAppid(),
		"toparty": 1,
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": "hello my dear friend",
		},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("msgid"))
	// Output: xxx
}

func ExampleRecallAgentMsg() {
	cli := getAgentClient()

	uri := "/cgi-bin/message/recall"
	params := map[string]interface{}{
		"msgid": "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}

func ExampleAgentCreateChat() {
	cli := getAgentClient()

	uri := "/cgi-bin/appchat/create"
	params := map[string]interface{}{
		"name":  "hello world",
		"owner": "xxx",
		"userlist": []string{
			"xxx",
			"xxxxxx",
		},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("chatid"))
	// Output: xxx
}

func ExampleAgentSendChatMsg() {
	cli := getAgentClient()

	uri := "/cgi-bin/appchat/send"
	params := map[string]interface{}{
		"chatid":  "xxx",
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": "welcome",
		},
		"safe": 1,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}
