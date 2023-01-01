package examples

import (
	"fmt"
	"log"
	"time"
)

func ExampleMicropay() {
	cli := getClient()

	authCode := "xxx"

	uri := "/pay/micropay"
	params := map[string]interface{}{
		"appid":            appid,
		"body":             "micropay test",
		"out_trade_no":     time.Now().Format("20060102150405"),
		"total_fee":        3,
		"spbill_create_ip": "127.0.0.1",
		"auth_code":        authCode,
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("transaction_id"))
	// Output: xxx
}
