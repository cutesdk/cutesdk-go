package examples

import (
	"fmt"
	"log"
)

func ExampleGetUserPhone() {
	cli := getClient()

	code := "xxx"

	uri := "/wxa/business/getuserphonenumber"
	params := map[string]interface{}{
		"code": code,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get user phone failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
