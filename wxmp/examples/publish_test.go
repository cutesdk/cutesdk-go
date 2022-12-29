package examples

import (
	"fmt"
	"log"
)

func ExampleGetPublishList() {
	client := getClient()

	uri := "/cgi-bin/freepublish/batchget"
	params := map[string]interface{}{
		"offset":     0,
		"count":      20,
		"no_content": 1,
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("send submsg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
