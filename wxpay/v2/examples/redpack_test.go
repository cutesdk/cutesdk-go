package examples

import (
	"fmt"
	"log"
	"time"
)

func ExampleSendRedpack() {
	cli := getClient()

	uri := "/mmpaymkttransfers/sendredpack"
	params := map[string]interface{}{
		"mch_billno":   time.Now().Format("20060102150405"),
		"wxappid":      "xxx",
		"send_name":    "redpack father",
		"re_openid":    "xxx",
		"total_amount": 300,
		"total_num":    1,
		"wishing":      "happy new year!",
		"client_ip":    "127.0.0.1",
		"act_name":     "new year redpack",
		"remark":       "happy new year with redpack",
		"scene_id":     "PRODUCT_1",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("send_listid"))

	// Output: xxx
}

func ExampleQueryRedpack() {
	cli := getClient()

	uri := "/mmpaymkttransfers/gethbinfo"
	params := map[string]interface{}{
		"appid":      appid,
		"mch_billno": "xxx",
		"bill_type":  "MCHT",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("detail_id"))
	// Output: xxx
}
