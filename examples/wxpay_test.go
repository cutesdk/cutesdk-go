package examples

import (
	"encoding/json"
	"fmt"

	"github.com/idoubi/cutesdk/wxpay"
	"github.com/idoubi/goutils"
)

// ExampleUnifiedOrder 统一下单示例
func ExampleUnifiedOrder() {
	sdk := getWxpaySDK()
	params := map[string]string{
		"appid": "wx25da2eca8fa3f4ef",
		"body": "JSAPI支付测试",
		"out_trade_no": goutils.GenOrderID(),
		"total_fee": "3",
		"notify_url": "https://idoubi.cc",
		"openid": "oLW495c2KVrduEpiSGDpHp7qKqCc",
	}

	res, err := sdk.UnifiedOrder(params)
	fmt.Printf("%s, %v\n", res, err)

	// Output: xxx
}

func ExampleGetPayParams() {
	sdk := getWxpaySDK()
	prepayID := "wx310014469480652fd422d88ff1a85e0000"
	appid := "wx25da2eca8fa3f4ef"

	res := sdk.GetPayParams(prepayID, appid)
	jb, _ := json.Marshal(&res)
	fmt.Printf("%s\n", jb)

	// Output: xxx
}

func getWxpaySDK() *wxpay.Wxpay {
	opts := wxpay.Options{
		Debug: true,
		MchID: "1498014222",
		APIKey: "Q5xQzBMvdvKQgn3Li3e26XVb4TNYuS13",
	}

	return wxpay.New(opts)
}
