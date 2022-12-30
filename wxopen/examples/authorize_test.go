package examples

import (
	"fmt"
	"log"
)

func ExampleGetAuthorizerList() {
	cli := getClient()

	uri := "/cgi-bin/component/api_get_authorizer_list"
	params := map[string]interface{}{
		"component_appid": cli.GetAppid(),
		"offset":          0,
		"count":           100,
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get authorizer list failed: %s\n", res)
	}

	fmt.Println(res.GetString("total_count"))
	// Output: xxx
}

func ExampleGetAuthorizerInfo() {
	cli := getClient()

	uri := "/cgi-bin/component/api_get_authorizer_info"
	params := map[string]interface{}{
		"component_appid":  cli.GetAppid(),
		"authorizer_appid": "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get authorizer info failed: %s\n", res)
	}

	fmt.Println(res.GetString("authorizer_info.nick_name"))
	// Output: xxx
}

func ExampleCreatePreAuthCode() {
	cli := getClient()

	uri := "/cgi-bin/component/api_create_preauthcode"
	params := map[string]interface{}{
		"component_appid": cli.GetAppid(),
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("create pre_auth_code failed: %s\n", res)
	}

	fmt.Println(res.GetString("pre_auth_code"))
	// Output: xxx
}

func ExampleGetPcAuthUrl() {
	cli := getClient()

	preAuthCode := "xxx"
	redirectUri := "https://xxx.com/wxopen/auth-callback"
	authType := "1"
	extra := map[string]string{
		"biz_appid":        "xxx",
		"category_id_list": "1|27",
	}

	authUrl, _ := cli.GetPcAuthUrl(preAuthCode, redirectUri, authType, extra)

	fmt.Println(authUrl)
	// Output: xxx
}

func ExampleGetH5AuthUrl() {
	cli := getClient()

	preAuthCode := "xxx"
	redirectUri := "https://xxx.com/wxopen/auth-callback"
	authType := "3"
	extra := map[string]string{}

	authUrl, _ := cli.GetH5AuthUrl(preAuthCode, redirectUri, authType, extra)

	fmt.Println(authUrl)
	// Output: xxx
}

func ExampleGetAuthorizerRefreshToken() {
	cli := getClient()

	uri := "/cgi-bin/component/api_query_auth"
	params := map[string]interface{}{
		"component_appid":    cli.GetAppid(),
		"authorization_code": "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get authorizer_refresh_token failed: %s\n", res)
	}

	fmt.Println(res.GetString("authorization_info.authorizer_refresh_token"))
	// Output: xxx
}

func ExampleGetAuthorizerAccessToken() {
	cli := getClient()

	uri := "/cgi-bin/component/api_authorizer_token"
	params := map[string]interface{}{
		"component_appid":          cli.GetAppid(),
		"authorizer_appid":         "xxx",
		"authorizer_refresh_token": "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get authorizer_access_token failed: %s\n", res)
	}

	fmt.Println(res.GetString("authorizer_access_token"))
	// Output: xxx
}
