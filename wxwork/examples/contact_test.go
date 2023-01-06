package examples

import (
	"fmt"
	"log"
)

func ExampleGetDepartmentList() {
	cli := getContactClient()

	uri := "/cgi-bin/department/simplelist"
	params := map[string]interface{}{
		"id": 0,
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("department"))
	// Output: xxx
}

func ExampleCreateDepartment() {
	cli := getContactClient()

	uri := "/cgi-bin/department/create"
	params := map[string]interface{}{
		"name":     "Developer Center",
		"name_en":  "dc",
		"parentid": 1,
		"order":    999,
		"id":       2,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetInt("id"))
	// Output: xxx
}

func ExampleGetUserList() {
	cli := getContactClient()

	uri := "/cgi-bin/user/list_id"
	params := map[string]interface{}{
		"cursor": "",
		"limit":  1000,
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("dept_user"))
	// Output: xxx
}

func ExampleUserOpenid() {
	cli := getContactClient()

	uri := "/cgi-bin/user/convert_to_openid"
	params := map[string]interface{}{
		"userid": "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("openid"))
	// Output: xxx
}

func ExampleGetJoinQrcode() {
	cli := getContactClient()

	uri := "/cgi-bin/corp/get_join_qrcode"
	params := map[string]interface{}{
		"size_type": 2,
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("join_qrcode"))
	// Output: xxx
}

func ExampleGetTagList() {
	cli := getContactClient()

	uri := "/cgi-bin/tag/list"

	res, err := cli.GetWithToken(uri)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("taglist"))
	// Output: xxx
}

func ExampleCreateTag() {
	cli := getContactClient()

	uri := "/cgi-bin/tag/create"
	params := map[string]interface{}{
		"tagname": "test",
		"tagid":   1,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetInt("tagid"))
	// Output: xxx
}

func ExampleAddTagUsers() {
	cli := getContactClient()

	uri := "/cgi-bin/tag/addtagusers"
	params := map[string]interface{}{
		"tagid":     1,
		"partylist": []int{1},
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetInt("errcode"))
	// Output: xxx
}

func ExampleGetTagUsers() {
	cli := getContactClient()

	uri := "/cgi-bin/tag/get"
	params := map[string]interface{}{
		"tagid": 1,
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetArray("userlist"))
	// Output: xxx
}
