package examples

import (
	"fmt"
	"log"
	"time"
)

func ExampleUnifiedOrder() {
	cli := getClient()

	uri := "/pay/unifiedorder"
	params := map[string]interface{}{
		"appid":            "xxx",
		"body":             "pay test",
		"out_trade_no":     time.Now().Format("20060102150405"),
		"total_fee":        3,
		"spbill_create_ip": "127.0.0.1",
		"notify_url":       "https://xxx.com/wxpay/v2/pay-notify/1498014222",
		"trade_type":       "JSAPI",
		"openid":           "xxx",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("prepay_id"))
	// Output: xxx
}

func ExampleGetPayParams() {
	cli := getClient()

	prepayId := "xxx"
	signType := "MD5"

	res, err := cli.GetPayParams(appid, prepayId, signType)
	if err != nil {
		log.Fatalf("get pay params failed: %v\n", err)
	}

	fmt.Println(res)
	// Output: xxx
}
