package examples

import (
	"fmt"
	"log"
)

func ExampleRegisterCompanyApp() {
	cli := getClient()

	uri := "/cgi-bin/component/fastregisterweapp?action=create"
	params := map[string]interface{}{
		"name":                 "xxx",
		"code":                 "xxx",
		"code_type":            1,
		"legal_persona_wechat": "xxx",
		"legal_persona_name":   "xxx",
		"component_phone":      "400-100-0000",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("register company app failed: %s\n", res)
	}

	fmt.Println(res.GetString("errmsg"))
	// Output: xxx
}

func ExampleRegisterPersonalApp() {
	cli := getClient()

	uri := "/wxa/component/fastregisterpersonalweapp?action=create"
	params := map[string]interface{}{
		"idname":          "xxx",
		"wxuser":          "xxx",
		"component_phone": "400-100-0000",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("register personal app failed: %s\n", res)
	}

	fmt.Println(res.GetString("authorize_url"))
	// Output: xxx
}

func ExampleQueryRegisterStatus() {
	cli := getClient()

	uri := "/wxa/component/fastregisterpersonalweapp?action=query"
	params := map[string]interface{}{
		"taskid": "xxx",
	}

	res, err := cli.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("query register status failed: %s\n", res)
	}

	fmt.Println(res.GetString("authorize_url"))
	// Output: xxx
}

func ExampleGetRegisterAuthUrl() {
	cli := getClient()

	appid := "xxx"
	redirectUri := "https://xxx.com/wxopen/register-auth"
	copyWxVerify := "1"

	authUrl, _ := cli.GetRegisterAuthUrl(appid, redirectUri, copyWxVerify)

	fmt.Println(authUrl)
	// Output: xxx
}

func ExampleGetRebindAuthUrl() {
	cli := getClient()

	appid := "xxx"
	redirectUri := "https://xxx.com/wxopen/rebind-auth"

	authUrl, _ := cli.GetRebindAuthUrl(appid, redirectUri)

	fmt.Println(authUrl)
	// Output: xxx
}
