package examples

import (
	"fmt"
	"log"
	"time"
)

func ExampleJsapiPay() {
	cli := getClient()

	uri := "/v3/pay/transactions/jsapi"
	params := map[string]interface{}{
		"appid":        appid,
		"mchid":        mchId,
		"description":  "test pay",
		"out_trade_no": time.Now().Format("20060102150405"),
		"notify_url":   fmt.Sprintf("https://xxx.com/wxpay/v3/pay-notify/%s", cli.GetMchId()),
		"amount": map[string]interface{}{
			"total": 3,
		},
		"payer": map[string]interface{}{
			"openid": "xxx",
		},
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("prepay_id"))
	// Output: xxx
}

func ExampleGetPayParams() {
	cli := getClient()

	prepayId := "xxx"
	signType := "RSA"

	res, err := cli.GetPayParams(appid, prepayId, signType)
	if err != nil {
		log.Fatalf("get pay params failed: %v\n", err)
	}

	fmt.Println(res)
	// Output: xxx
}
