package examples

import (
	"fmt"
	"log"
)

func ExampleGetCustomerList() {
	cli := getCustomerClient()

	uri := "/cgi-bin/externalcontact/list"
	params := map[string]interface{}{
		"userid": "xxx",
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("external_userid"))
	// Output: xxx
}

func ExampleGetCustomers() {
	cli := getCustomerClient()

	uri := "/cgi-bin/externalcontact/batch/get_by_user"
	params := map[string]interface{}{
		"userid_list": []string{"xxx", "xxxxxx"},
		"cursor":      "",
		"limit":       100,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("external_contact_list"))
	// Output: xxx
}

func ExampleGetCustomerInfo() {
	cli := getCustomerClient()

	uri := "/cgi-bin/externalcontact/get"
	params := map[string]interface{}{
		"external_userid": "xxx",
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("external_contact.unionid"))
	// Output: xxx
}

func ExampleGetCustomerGroups() {
	cli := getCustomerClient()

	uri := "/cgi-bin/externalcontact/groupchat/list"
	params := map[string]interface{}{
		"status_filter": 0,
		"owner_filter": map[string]interface{}{
			"userid_list": []string{"idoubi"},
		},
		"cursor": "",
		"limit":  1000,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("group_chat_list"))
	// Output: xxx
}

func ExampleGetCustomerGroup() {
	cli := getCustomerClient()

	uri := "/cgi-bin/externalcontact/groupchat/get"
	params := map[string]interface{}{
		"chat_id":   "xxx",
		"need_name": 1,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("group_chat.member_list"))
	// Output: xxx
}

func ExampleTransferCustomer() {
	cli := getCustomerClient()

	uri := "/cgi-bin/externalcontact/transfer_customer"
	params := map[string]interface{}{
		"handover_userid": "xxx",
		"takeover_userid": "xxxxxx",
		"external_userid": []string{
			"xxx",
		},
		"transfer_success_msg": "hello, my friend",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("customer"))
	// Output: xxx
}

func ExampleSendWelcomeMsg() {
	cli := getCustomerClient()

	uri := "/cgi-bin/externalcontact/send_welcome_msg"
	params := map[string]interface{}{
		"welcome_code": "xxx",
		"text": map[string]interface{}{
			"content": "welcome my friend",
		},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetInt("errcode"))
	// Output: xxx
}
