package examples

import (
	"fmt"
	"log"
	"time"
)

func ExampleSendSubmsg() {
	client := getClient()

	uri := "/cgi-bin/message/subscribe/bizsend"
	params := map[string]interface{}{
		"touser":      "xxx",
		"template_id": "xxx",
		"page":        "https://baidu.com",
		"miniprogram": map[string]interface{}{
			"appid":    "xxx",
			"pagepath": "pages/index/index",
		},
		"data": map[string]interface{}{
			"thing5": map[string]interface{}{
				"value": "user name",
			},
			"thing7": map[string]interface{}{
				"value": "message content",
			},
			"time2": map[string]interface{}{
				"value": time.Now().Format("2006-01-02 15:04:05"),
			},
		},
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
