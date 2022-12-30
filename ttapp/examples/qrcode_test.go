package examples

import (
	"fmt"
	"log"
)

func ExampleCreateQrcode() {
	cli := getClient()

	uri := "/api/apps/qrcode"
	params := map[string]interface{}{
		"appname": "douyin",
		"path":    "",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") > 0 {
		log.Fatalf("create qrcode failed: %s\n", res)
	}

	if err := res.SaveAsFile("./qrcode.jpg"); err != nil {
		log.Fatalf("save qrcode failed: %v\n", err)
	}

	fmt.Println("ok")
	// Output: ok
}
