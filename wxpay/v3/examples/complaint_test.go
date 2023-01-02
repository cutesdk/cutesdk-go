package examples

import (
	"fmt"
	"log"
)

func ExampleGetComplaints() {
	cli := getClient()

	beginDate := "2022-10-01"
	endDate := "2022-10-30"

	uri := "/v3/merchant-service/complaints-v2"
	params := map[string]interface{}{
		"begin_date": beginDate,
		"end_date":   endDate,
	}

	res, err := cli.Get(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetInt("total_count"))
	// Output: xxx
}

func ExampleGetComplaintInfo() {
	cli := getClient()

	complaintId := "xxx"

	uri := fmt.Sprintf("/v3/merchant-service/complaints-v2/%s", complaintId)

	res, err := cli.Get(uri)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("complaint_state"))
	// Output: xxx
}

func ExampleGetComplaintNotifyUrl() {
	cli := getClient()

	uri := "/v3/merchant-service/complaint-notifications"

	res, err := cli.Get(uri)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("url"))
	// Output: xxx
}

func ExampleCreateComplaintNotifyUrl() {
	cli := getClient()

	uri := "/v3/merchant-service/complaint-notifications"

	params := map[string]interface{}{
		"url": fmt.Sprintf("https://xxx.com/complaint/notify/%s", cli.GetMchId()),
	}

	res, err := cli.Post(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("url"))
	// Output: xxx
}

func ExampleUpdateComplaintNotifyUrl() {
	cli := getClient()

	uri := "/v3/merchant-service/complaint-notifications"

	params := map[string]interface{}{
		"url": fmt.Sprintf("https://xxx.com/complaint/notify/%s", cli.GetMchId()),
	}

	res, err := cli.Put(uri, params)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("url"))
	// Output: xxx
}

func ExampleDeleteComplaintNotifyUrl() {
	cli := getClient()

	uri := "/v3/merchant-service/complaint-notifications"

	res, err := cli.Delete(uri)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %s, %s\n", wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res)
	// Output: xxxx
}
