package examples

import (
	"fmt"
	"log"
)

func ExampleGetOauthUrl() {
	cli := getClient()

	redirectUri := "https://xxx.com/api/callback/"
	scope := "snsapi_userinfo"
	state := "test123"
	extra := map[string]string{
		"forcePopup":    "false",
		"forceSnapShot": "false",
	}

	oauthUrl, err := cli.GetOauthUrl(redirectUri, scope, state, extra)
	if err != nil {
		log.Fatalf("get oauth url failed: %v\n", err)
	}

	fmt.Println(oauthUrl)
	// Output: xxx
}

func ExampleGetOauthToken() {
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

func ExampleGetOauthUser() {
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
