package examples

import (
	"fmt"
	"log"
)

func ExampleSendSubmsg() {
	client := getClient()

	uri := "/cgi-bin/message/subscribe/send"
	params := map[string]interface{}{
		"template_id": "xxx",
		"page":        "pages/index/index",
		"touser":      "xxx",
		"data": map[string]interface{}{
			"name1": map[string]interface{}{
				"value": "message title",
			},
			"thing3": map[string]interface{}{
				"value": "message content",
			},
		},
		"miniprogram_state": "develop",
		"lang":              "zh_CN",
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("send submsg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
