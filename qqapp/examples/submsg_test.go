package examples

import (
	"fmt"
	"log"
)

func ExampleSendSubmsg() {
	cli := getClient()

	uri := "/api/json/subscribe/SendSubscriptionMessage"
	params := map[string]interface{}{
		"touser":      "xxx",
		"template_id": "xxx",
		"data": map[string]interface{}{
			"keyword1": map[string]interface{}{
				"value": "message title",
			},
			"keyword2": map[string]interface{}{
				"value": "message content",
			},
		},
		"emphasis_keyword": "keyword1.DATA",
		"page":             "pages/index/index",
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
