# qqapp

qq miniapp sdk

## install

```shell
go get -u github.com/cutesdk/cutesdk-go/qqapp
```

## usage

take [create qrcode api](https://q.qq.com/wiki/develop/game/server/open-port/qr-code.html) for example: 

```go
package main

import (
	"log"

	"github.com/cutesdk/cutesdk-go/qqapp"
)

func main() {
	// new qqapp client
	client, err := qqapp.NewClient(&qqapp.Options{
		Debug:  true,
		Appid:  "xxx",
		Secret: "xxx",
	})
	if err != nil {
		log.Fatalf("new qqapp client failed: %v", err)
	}

	// request create qrcode api
	uri := "/api/json/qqa/CreateMiniCode"
	params := map[string]interface{}{
		"appid": client.GetAppid(),
		"path":  "/pages/index/index",
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