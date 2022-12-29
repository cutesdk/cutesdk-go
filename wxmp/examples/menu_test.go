package examples

import (
	"fmt"
	"log"
)

func ExampleGetCurrentMenu() {
	client := getClient()

	uri := "/cgi-bin/get_current_selfmenu_info"

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
