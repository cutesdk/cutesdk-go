package examples

import (
	"fmt"
	"log"
)

func ExampleCode2Session() {
	client := getClient()

	code := "xxx"
	anonymousCode := "xxx"

	uri := "/api/apps/v2/jscode2session"
	params := map[string]interface{}{
		"appid":          client.GetAppid(),
		"secret":         client.GetSecret(),
		"code":           code,
		"anonymous_code": anonymousCode,
	}

	res, err := client.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	openid := res.GetString("data.openid")
	if openid == "" {
		log.Fatalf("code2session failed: %s\n", res)
	}

	fmt.Println(openid)
	// Output: xxx
}