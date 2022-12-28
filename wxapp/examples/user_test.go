package examples

import (
	"fmt"
	"log"
)

func ExampleGetUserPhone() {
	client := getClient()

	code := "xxx"

	uri := "/wxa/business/getuserphonenumber"
	params := map[string]interface{}{
		"code": code,
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get user phone failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
