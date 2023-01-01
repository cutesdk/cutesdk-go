package examples

import (
	"fmt"
	"log"
)

func ExampleQueryByOutTradeNo() {
	cli := getClient()

	uri := "/pay/orderquery"
	params := map[string]interface{}{
		"appid":        appid,
		"out_trade_no": "xxx",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("transaction_id"))
	// Output: xxx
}

func ExampleQueryByTransactionId() {
	cli := getClient()

	uri := "/pay/orderquery"
	params := map[string]interface{}{
		"appid":          appid,
		"transaction_id": "xxx",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("out_trade_no"))
	// Output: xxx
}
