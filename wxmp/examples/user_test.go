package examples

import (
	"fmt"
	"log"
)

func ExampleGetTags() {
	cli := getClient()

	uri := "/cgi-bin/tags/get"

	res, err := cli.GetWithToken(uri)
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
	cli := getClient()

	uri := "/cgi-bin/user/get"
	params := map[string]interface{}{
		"next_openid": "",
	}

	res, err := cli.GetWithToken(uri, params)
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
	cli := getClient()

	uri := "/cgi-bin/user/info"
	params := map[string]interface{}{
		"openid": "xxx",
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get tags failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
