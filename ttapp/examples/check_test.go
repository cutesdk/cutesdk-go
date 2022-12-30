package examples

import (
	"fmt"
	"log"
)

func ExampleCheckMsg() {
	cli := getClient()

	accessToken, err := cli.GetAccessToken()
	if err != nil {
		log.Fatalf("get access_token failed: %v\n", err)
	}

	uri := "/api/v2/tags/text/antidirt"
	params := map[string]interface{}{
		"tasks": []map[string]interface{}{
			{"content": "hello"},
		},
	}
	headers := map[string]interface{}{
		"X-Token": accessToken,
	}

	res, err := cli.Post(uri, params, headers)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("code") != 0 {
		log.Fatalf("check msg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
