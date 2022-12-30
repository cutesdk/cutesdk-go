package examples

import (
	"fmt"
	"log"
)

func ExampleCheckMsg() {
	cli := getClient()

	uri := "/wxa/msg_sec_check"
	params := map[string]interface{}{
		"content": "hello",
		"version": 2,
		"scene":   2,
		"openid":  "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errCode") != 0 {
		log.Fatalf("check msg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
