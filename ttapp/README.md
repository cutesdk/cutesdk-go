# ttapp

bytedance miniapp sdk

## usage

take [create qrcode api](https://microapp.bytedance.com/docs/zh-CN/mini-app/develop/server/qr-code/create-qr-code) for example: 

```go
package main

import (
	"log"

	"github.com/cutesdk/cutesdk-go/ttapp"
)

func main() {
	// new ttapp client
	client, err := ttapp.NewClient(&ttapp.Options{
		Debug:  true,
		Appid:  "xxx",
		Secret: "xxx",
	})
	if err != nil {
		log.Fatalf("new ttapp client failed: %v", err)
	}

	// request create qrcode api
	uri := "/api/apps/qrcode"
	params := map[string]interface{}{
		"appname": "douyin",
		"path":    "",
	}

	// auto fetch access_token
	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	// deal with api response
	if res.GetInt("errcode") > 0 {
		log.Fatalf("create qrcode failed: %s\n", res)
	}

	if err := res.SaveAsFile("./qrcode.jpg"); err != nil {
		log.Fatalf("save qrcode failed: %v\n", err)
	}

	log.Println("ok")
}
```

## more

view some other api demos in `examples` folder.