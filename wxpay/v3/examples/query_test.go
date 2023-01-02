package examples

import (
	"fmt"
	"log"
)

func ExampleQueryByOutTradeNo() {
	cli := getClient()

	outTradeNo := "xxx"

	uri := fmt.Sprintf("/v3/pay/transactions/out-trade-no/%s", outTradeNo)
	params := map[string]interface{}{
		"mchid": mchId,
	}

	res, err := cli.Get(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %d, %s, %s\n", wxerr.StatusCode, wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("transaction_id"))
	// Output: xxx
}

func ExampleQueryByTransactionId() {
	cli := getClient()

	transactionId := "xxx"

	uri := fmt.Sprintf("/v3/pay/transactions/id/%s", transactionId)
	params := map[string]interface{}{
		"mchid": mchId,
	}

	res, err := cli.Get(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %d, %s, %s\n", wxerr.StatusCode, wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("out_trade_no"))
	// Output: xxx
}
