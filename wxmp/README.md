# wxmp

weixin mp sdk

## usage

take [create qrcode api](https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html) for example: 

```go
package main

import (
	"log"

	"github.com/cutesdk/cutesdk-go/wxmp"
)

func main() {
	// new wxmp client
	client, err := wxmp.NewClient(&wxmp.Options{
		Debug:  true,
		Appid:  "xxx",
		Secret: "xxx",
	})
	if err != nil {
		log.Fatalf("new wxmp client failed: %v", err)
	}

	// request create qrcode api
	uri := "/cgi-bin/qrcode/create"
	params := map[string]interface{}{
		"expire_seconds": 300,
		"action_name":    "QR_STR_SCENE",
		"action_info": map[string]interface{}{
			"scene": map[string]interface{}{
				"scene_str": "id=3",
			},
		},
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

	ticket := res.GetString("ticket")
	picurl := fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s", ticket)
	log.Printf("picurl: %s\n", picurl)
}
```

## more

view some other api demos in `examples` folder.