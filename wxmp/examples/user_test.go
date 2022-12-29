package examples

import (
	"fmt"
	"log"
)

func ExampleGetTags() {
	client := getClient()

	uri := "/cgi-bin/tags/get"

	res, err := client.GetWithToken(uri)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get tags failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleGetUserOpenids() {
	client := getClient()

	uri := "/cgi-bin/user/get"
	params := map[string]interface{}{
		"next_openid": "",
	}

	res, err := client.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get tags failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleGetUserInfo() {
	client := getClient()

	uri := "/cgi-bin/user/info"
	params := map[string]interface{}{
		"openid": "xxx",
	}

	res, err := client.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get tags failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
