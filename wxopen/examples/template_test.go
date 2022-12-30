package examples

import (
	"fmt"
	"log"
)

func ExampleGetTemplateList() {
	cli := getClient()

	uri := "/wxa/gettemplatelist"
	params := map[string]interface{}{
		"template_type": 0,
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get template list failed: %s\n", res)
	}

	fmt.Println(res.GetArray("template_list"))
	// Output: xxx
}

func ExampleGetTemplateDraftList() {
	cli := getClient()

	uri := "/wxa/gettemplatedraftlist"

	res, err := cli.GetWithToken(uri)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get template draft list failed: %s\n", res)
	}

	fmt.Println(res.GetArray("draft_list"))
	// Output: xxx
}
