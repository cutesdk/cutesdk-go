package examples

import (
	"encoding/json"
	"fmt"
	"log"
)

func ExampleGetJssdkConfig() {
	cli := getClient()

	url := "https://xxx.com?p=123"

	res, err := cli.GetJssdkConfig(url)

	if err != nil {
		log.Fatalf("get jssdk config failed: %v", err)
	}

	j, _ := json.Marshal(res)

	fmt.Printf("%s", j)
	// Output: xxx
}
