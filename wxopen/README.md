# wxopen

weixin open platform sdk

## usage

take [register personal app api](https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/register-management/fast-registration-ind/fastRegisterPersonalMp.html) for example: 

```go
package main

import (
	"log"

	"github.com/cutesdk/cutesdk-go/wxopen"
)

func main() {
	// new wxopen client
	cli, err := wxopen.NewClient(&wxopen.Options{
		Debug:  true,
		Appid:  "xxx",
		Secret: "xxx",
	})
	if err != nil {
		log.Fatalf("new wxopen client failed: %v", err)
	}

	// request create qrcode api
	uri := "/wxa/component/fastregisterpersonalweapp?action=create"
	params := map[string]interface{}{
		"idname":          "xxx",
		"wxuser":          "xxx",
		"component_phone": "400-100-0000",
	}

	// auto fetch access_token
	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	// deal with api response
	if res.GetInt("errcode") > 0 {
		log.Fatalf("create qrcode failed: %s\n", res)
	}

	log.Println("ok")
}
```

## more

view some other api demos in `examples` folder.