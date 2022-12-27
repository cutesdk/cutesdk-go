package examples

import (
	"fmt"
	"log"
)

func ExampleCheckMsg() {
	client := getClient()

	uri := "/api/json/security/MsgSecCheck"
	params := map[string]interface{}{
		"appid":   client.GetAppid(),
		"content": "hello",
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errCode") != 0 {
		log.Fatalf("check msg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
