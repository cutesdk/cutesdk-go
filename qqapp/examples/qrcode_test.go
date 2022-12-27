package examples

import (
	"fmt"
	"log"
)

func ExampleCreateQrcode() {
	client := getClient()

	uri := "/api/json/qqa/CreateMiniCode"
	params := map[string]interface{}{
		"appid": client.GetAppid(),
		"path":  "pages/index/index",
	}

	res, err := client.PostWithToken(uri, params)
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
