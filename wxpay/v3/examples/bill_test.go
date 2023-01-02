package examples

import (
	"fmt"
	"log"
)

func ExampleApplyPayBill() {
	cli := getClient()

	uri := "/v3/bill/tradebill"
	params := map[string]interface{}{
		"bill_date": "2022-12-30",
		"bill_type": "ALL",
		"tar_type":  "GZIP",
	}

	res, err := cli.Get(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	downloadUrl := res.GetString("download_url")

	bill, err := cli.WithoutValidator().Get(downloadUrl)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("download failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("download failed: %v\n", err)
	}

	if err := bill.SaveAsFile("./pay_bill.gzip"); err != nil {
		log.Fatalf("save file failed: %v\n", err)
	}

	fmt.Println("xxx")
	// Output: xxx
}

func ExampleApplyFundBill() {
	cli := getClient()

	uri := "/v3/bill/fundflowbill"
	params := map[string]interface{}{
		"bill_date":    "2022-12-30",
		"account_type": "BASIC",
	}

	res, err := cli.Get(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	downloadUrl := res.GetString("download_url")

	bill, err := cli.WithoutValidator().Get(downloadUrl)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("download failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("download failed: %v\n", err)
	}

	if err := bill.SaveAsFile("./fund_bill.xls"); err != nil {
		log.Fatalf("save file failed: %v\n", err)
	}

	fmt.Println("xxx")
	// Output: xxx
}
