package examples

import (
	"fmt"
	"log"
)

func ExampleGetOauthUrl() {
	client := getClient()

	redirectUri := "https://www.xxx.com/api/callback/"
	scope := "snsapi_userinfo"
	state := "test123"

	oauthUrl, err := client.GetOauthUrl(redirectUri, scope, state)
	if err != nil {
		log.Fatalf("get oauth url failed: %v\n", err)
	}

	fmt.Println(oauthUrl)
	// Output: xxx
}

func ExampleGetOauthToken() {
	client := getClient()

	code := "xxx"

	res, err := client.GetOauthToken(code)
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
	client := getClient()

	oauthAccessToken := "xxx"
	openid := "xxx"

	res, err := client.GetOauthUser(oauthAccessToken, openid)
	if err != nil {
		log.Fatalf("get oauth user failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get oauth user failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
