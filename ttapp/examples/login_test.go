package examples

import (
	"fmt"
	"log"
)

func ExampleCode2Session() {
	cli := getClient()

	code := "xxx"
	anonymousCode := "xxx"

	res, err := cli.Login(code, anonymousCode)
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
