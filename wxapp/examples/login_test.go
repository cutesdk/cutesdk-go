package examples

import (
	"fmt"
	"log"
)

func ExampleCode2Session() {
	client := getClient()

	code := "xxx"

	uri := "/sns/jscode2session"
	params := map[string]interface{}{
		"appid":      client.GetAppid(),
		"secret":     client.GetSecret(),
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	res, err := client.Get(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	openid := res.GetString("openid")
	if openid == "" {
		log.Fatalf("code2session failed: %s\n", res)
	}

	fmt.Println(openid)
	// Output: xxx
}
