package examples

import (
	"fmt"
	"log"
)

func ExampleWxopenGetOauthUrl() {
	cli := getClient()

	redirectUri := "https://xxx.com/wxopen/login/callback/"
	scope := "snsapi_login"
	extra := map[string]string{
		"state": "test123",
	}

	oauthUrl, err := cli.GetOauthUrl(redirectUri, scope, extra)
	if err != nil {
		log.Fatalf("get oauth url failed: %v\n", err)
	}

	fmt.Println(oauthUrl)
	// Output: xxx
}

func ExampleWxopenGetOauthToken() {
	cli := getClient()

	code := "xxx"

	res, err := cli.GetOauthToken(code)
	if err != nil {
		log.Fatalf("get oauth token failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get oauth token failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleWxopenGetOauthUser() {
	cli := getClient()

	oauthAccessToken := "xxx"
	openid := "xxx"

	res, err := cli.GetOauthUser(oauthAccessToken, openid)
	if err != nil {
		log.Fatalf("get oauth user failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get oauth user failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
