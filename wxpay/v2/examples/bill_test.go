package examples

import (
	"fmt"
	"log"
)

func ExampleDownloadPayBill() {
	cli := getClient()

	uri := "/pay/downloadbill"
	params := map[string]interface{}{
		"appid":     appid,
		"bill_date": "20221230",
		"bill_type": "ALL",
		"tar_type":  "GZIP",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if errCode := res.GetString("error_code"); errCode != "" {
		log.Fatalf("download pay bill failed: %s\n", res)
	}

	if err := res.SaveAsFile("./pay_bill.gzip"); err != nil {
		log.Fatalf("save file failed: %v\n", err)
	}

	fmt.Println("xxx")
	// Output: xxx
}

func ExampleDownloadFundBill() {
	cli := getClient()

	uri := "/pay/downloadfundflow"
	params := map[string]interface{}{
		"appid":        appid,
		"bill_date":    "20221230",
		"account_type": "Basic",
		"sign_type":    "HMAC-SHA256",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if errCode := res.GetString("error_code"); errCode != "" {
		log.Fatalf("download pay bill failed: %s\n", res)
	}

	if err := res.SaveAsFile("./fund_bill.xls"); err != nil {
		log.Fatalf("save file failed: %v\n", err)
	}

	fmt.Println("xxx")
	// Output: xxx
}
