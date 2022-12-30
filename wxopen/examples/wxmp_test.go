package examples

import (
	"encoding/json"
	"fmt"
	"log"
)

func ExampleGetBindWxapps() {
	wxmpCli := getWxmpClient()

	uri := "/cgi-bin/wxopen/wxamplinkget"

	res, err := wxmpCli.PostWithToken(uri)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get bind wxapps failed: %s\n", res)
	}

	fmt.Println(res.GetArray("wxopens.items"))
	// Output: xxx
}

func ExampleGetJssdkConfig() {
	wxmpCli := getWxmpClient()

	url := "https://xxx.com?p=123"

	res, err := wxmpCli.GetJssdkConfig(url)

	if err != nil {
		log.Fatalf("get jssdk config failed: %v", err)
	}

	j, _ := json.Marshal(res)

	fmt.Printf("%s", j)
	// Output: xxx
}

func ExampleGetOauthUrl() {
	cli := getClient()
	wxmpCli := getWxmpClient()

	redirectUri := "https://xxx.com/api/callback/"
	scope := "snsapi_userinfo"
	state := "test123"
	extra := map[string]string{
		"forcePopup":      "false",
		"forceSnapShot":   "false",
		"component_appid": cli.GetAppid(),
	}

	oauthUrl, err := wxmpCli.GetOauthUrl(redirectUri, scope, state, extra)
	if err != nil {
		log.Fatalf("get oauth url failed: %v\n", err)
	}

	fmt.Println(oauthUrl)
	// Output: xxx
}

func ExampleGetOauthToken() {
	wxmpCli := getWxmpClient()

	code := "xxx"

	res, err := wxmpCli.GetOauthToken(code)
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
	wxmpCli := getWxmpClient()

	oauthAccessToken := "xxx"
	openid := "xxx"

	res, err := wxmpCli.GetOauthUser(oauthAccessToken, openid)
	if err != nil {
		log.Fatalf("get oauth user failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get oauth user failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
