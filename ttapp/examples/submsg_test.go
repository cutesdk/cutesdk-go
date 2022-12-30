package examples

import (
	"fmt"
	"log"
)

func ExampleSendSubmsg() {
	cli := getClient()

	uri := "/api/apps/subscribe_notification/developer/v1/notify"
	params := map[string]interface{}{
		"app_id":  cli.GetAppid(),
		"tpl_id":  "xxx",
		"open_id": "xxx",
		"data": map[string]interface{}{
			"留言人":  "idoubi",
			"留言内容": "嘻嘻😄",
		},
		"page": "pages/index/index",
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
