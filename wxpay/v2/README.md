# wxpay v2

wxpay v2 sdk

## usage

take [unifiedorder api](https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1) for example: 

```go
package main

import (
	"log"

	"github.com/cutesdk/cutesdk-go/wxpay/v2"
)

func main() {
	// new wxpay client
	cli, err := wxpay.NewClient(&wxpay.Options{
		MchId:    "xxx",
		ApiKey:   "xxx",
		CertFile: "./apiclient_cert.pem",
		KeyFile:  "./apiclient_key.pem",
		Debug:    true,
	})
	if err != nil {
		log.Fatalf("new wxpay client failed: %v", err)
	}

	// request unifiedorder api
	uri := "/pay/unifiedorder"
	params := map[string]interface{}{
		"appid":            "xxx",
		"body":             "pay test",
		"out_trade_no":     "xxx",
		"total_fee":        3,
		"spbill_create_ip": "127.0.0.1",
		"notify_url":       "https://xxx.com/wxpay/v2/pay-notify/1498014222",
		"trade_type":       "JSAPI",
		"openid":           "xxx",
	}

	// auto gen sign
	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	// deal with api response
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	log.Println(res.GetString("prepay_id"))
}
```

## more

view some other api demos in `examples` folder.