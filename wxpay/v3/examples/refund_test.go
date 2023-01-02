package examples

import (
	"fmt"
	"log"
)

func ExampleCloseOrder() {
	cli := getClient()

	outTradeNo := "xxx"

	uri := fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s/close", outTradeNo)
	params := map[string]interface{}{
		"mchid": cli.GetMchId(),
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %d, %s, %s\n", wxerr.StatusCode, wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res)
	// Output: xxx
}

func ExampleRefund() {
	cli := getClient()

	uri := "/v3/refund/domestic/refunds"
	params := map[string]interface{}{
		"out_trade_no":  "xxx",
		"out_refund_no": "xxx_RF_02",
		"amount": map[string]interface{}{
			"refund":   1,
			"total":    3,
			"currency": "CNY",
		},
		"notify_url": fmt.Sprintf("https://idoubi.cc/wxpay/v3/refund-notify/%s", cli.GetMchId()),
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %d, %s, %s\n", wxerr.StatusCode, wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("refund_id"))
	// Output: xxx
}

func ExampleQueryRefundByOutRefundNo() {
	cli := getClient()

	outRefundNo := "xxx_RF_01"

	uri := fmt.Sprintf("/v3/refund/domestic/refunds/%s", outRefundNo)

	res, err := cli.Get(uri)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("refund_id"))
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
