package examples

import (
	"fmt"
	"log"
)

func ExampleCreateQrcode() {
	cli := getClient()

	uri := "/wxa/getwxacodeunlimit"
	params := map[string]interface{}{
		"scene":       "id=3",
		"page":        "pages/index/index",
		"check_path":  false,
		"env_version": "develop",
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
