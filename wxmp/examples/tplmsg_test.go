package examples

import (
	"fmt"
	"log"
)

func ExampleSendTplmsg() {
	cli := getClient()

	uri := "/cgi-bin/message/template/send"
	params := map[string]interface{}{
		"touser":        "xxx",
		"template_id":   "xxx",
		"url":           "https://baidu.com",
		"client_msg_id": "xxx",
		"data": map[string]interface{}{
			"first": map[string]interface{}{
				"value": "message title",
				"color": "#08a5e0",
			},
			"user": map[string]interface{}{
				"value": "user name",
			},
			"ask": map[string]interface{}{
				"value": "ask content",
			},
			"remark": map[string]interface{}{
				"value": "remark info",
				"color": "#1ba358",
			},
		},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("send submsg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
