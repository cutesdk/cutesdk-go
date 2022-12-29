package examples

import (
	"encoding/json"
	"fmt"
	"log"
)

func ExampleGetJssdkConfig() {
	client := getClient()

	url := "https://xxx.com?p=123"

	res, err := client.GetJssdkConfig(url)

	if err != nil {
		log.Fatalf("get jssdk config failed: %v", err)
	}

	j, _ := json.Marshal(res)

	fmt.Printf("%s", j)
	// Output: xxx
}
