package examples

import (
	"fmt"
	"log"
)

func ExampleCloseOrder() {
	cli := getClient()

	uri := "/pay/closeorder"
	params := map[string]interface{}{
		"appid":        appid,
		"out_trade_no": "xxx",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("result_code"))
	// Output: xxx
}

func ExampleRefund() {
	cli := getClient()

	uri := "/secapi/pay/refund"
	params := map[string]interface{}{
		"appid":         appid,
		"out_trade_no":  "xxx",
		"out_refund_no": "xxx_RF_01",
		"total_fee":     3,
		"refund_fee":    1,
		"notify_url":    "https://xxx.com/wxpay/v2/refund-notify/1498014222",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("refund_id"))
	// Output: xxx
}

func ExampleQueryRefundByOutRefundNo() {
	cli := getClient()

	outRefundNo := "xxx"

	uri := "/pay/refundquery"
	params := map[string]interface{}{
		"appid":         appid,
		"out_refund_no": outRefundNo,
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("out_trade_no"))
	// Output: xxx
}

func ExampleQueryRefundByRefundId() {
	cli := getClient()

	refundId := "xxx"

	uri := "/pay/refundquery"
	params := map[string]interface{}{
		"appid":     appid,
		"refund_id": refundId,
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("out_trade_no"))
	// Output: xxx
}

func ExampleQueryRefundByOutTradeNo() {
	cli := getClient()

	outTradeNo := "xxx"

	uri := "/pay/refundquery"
	params := map[string]interface{}{
		"appid":        appid,
		"out_trade_no": outTradeNo,
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("out_trade_no"))
	// Output: xxx
}
