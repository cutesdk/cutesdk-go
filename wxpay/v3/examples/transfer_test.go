package examples

import (
	"fmt"
	"log"
	"time"
)

func ExampleTransfer() {
	cli := getClient()

	batchid := time.Now().Format("20060102150405")

	uri := "/v3/transfer/batches"
	params := map[string]interface{}{
		"appid":        "xxx",
		"out_batch_no": batchid,
		"batch_name":   "test batch name",
		"batch_remark": "test batch remark",
		"total_amount": 300,
		"total_num":    1,
		"transfer_detail_list": []map[string]interface{}{
			{
				"out_detail_no":   fmt.Sprintf("%s01", batchid),
				"transfer_amount": 300,
				"transfer_remark": "test transfer remark",
				"openid":          "xxx",
			},
		},
		"transfer_scene_id": "1001",
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("batch_id"))
	// Output: xxx
}

func ExampleQueryTransfer() {
	cli := getClient()

	outBatchNo := "xxx"
	uri := fmt.Sprintf("/v3/transfer/batches/out-batch-no/%s", outBatchNo)
	params := map[string]interface{}{
		"need_query_detail": true,
		"offset":            0,
		"limit":             20,
		"detail_status":     "ALL",
	}

	res, err := cli.Get(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("transfer_batch.batch_id"))
	// Output: xxx
}
