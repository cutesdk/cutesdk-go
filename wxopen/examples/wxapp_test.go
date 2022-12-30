package examples

import (
	"fmt"
	"log"
)

func ExampleWxappLogin() {
	wxappCli := getWxappClient()

	code := "xxx"

	res, err := wxappCli.Login(code)

	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("wxapp login failed: %s\n", res)
	}

	fmt.Println(res.GetString("openid"))
	// Output: xxx
}

func ExampleGetWxappInfo() {
	wxappCli := getWxappClient()

	uri := "/cgi-bin/account/getaccountbasicinfo"

	res, err := wxappCli.GetWithToken(uri)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get wxapp account info failed: %s\n", res)
	}

	fmt.Println(res.GetString("nickname"))
	// Output: xxx
}

func ExampleVerifyWxappName() {
	wxappCli := getWxappClient()

	uri := "/cgi-bin/wxverify/checkwxverifynickname"
	params := map[string]interface{}{
		"nick_name": "xxx",
	}

	res, err := wxappCli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("check wxapp name failed: %s\n", res)
	}

	fmt.Println(res.GetString("wording"))
	// Output: xxx
}

func ExampleGetMemberList() {
	wxappCli := getWxappClient()

	uri := "/wxa/memberauth"
	params := map[string]interface{}{
		"action": "get_experiencer",
	}

	res, err := wxappCli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("check wxapp name failed: %s\n", res)
	}

	fmt.Println(res.GetArray("members"))
	// Output: xxx
}
