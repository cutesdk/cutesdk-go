package examples

import (
	"fmt"
	"log"
)

func ExampleCheckMsg() {
	cli := getClient()

	uri := "/api/json/security/MsgSecCheck"
	params := map[string]interface{}{
		"appid":   cli.GetAppid(),
		"content": "hello",
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
