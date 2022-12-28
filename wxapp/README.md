# wxapp

weixin miniapp sdk

## usage

take [create qrcode api](https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/qr-code/getUnlimitedQRCode.html) for example: 

```go
package main

import (
	"log"

	"github.com/cutesdk/cutesdk-go/wxapp"
)

func main() {
	// new wxapp client
	client, err := wxapp.NewClient(&wxapp.Options{
		Debug:  true,
		Appid:  "xxx",
		Secret: "xxx",
	})
	if err != nil {
		log.Fatalf("new wxapp client failed: %v", err)
	}

	// request create qrcode api
	uri := "/wxa/getwxacodeunlimit"
	params := map[string]interface{}{
		"scene":       "id=3",
		"page":        "pages/index/index",
		"check_path":  false,
		"env_version": "develop",
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