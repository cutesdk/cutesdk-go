package examples

import (
	"fmt"

	"github.com/idoubi/cutesdk/wxapp"
)

// ExampleCode2Session 小程序登录示例
func ExampleCode2Session() {
	sdk := getWxappSdk()

	res, err := sdk.Get("/sns/jscode2session", map[string]string{
		"appid":      "wx25da2eca8fa3f4ef",
		"secret":     "6cca21a5d4b694574fa912a0a928a742",
		"js_code":    "xx",
		"grant_type": "authorization_code",
	})

	fmt.Println(res, err)
	// Output: xxx
}

// ExampleApiGet 请求GET类型API
func ExampleApiGet() {
	sdk := getWxappSdk()

	res, err := sdk.Get("/wxa/getnearbypoilist", map[string]string{
		"page":      "1",
		"page_rows": "200",
	})

	fmt.Println(res.Get("data").String(), err)

	// Output: xxx
}

// ExampleApiPost 请求POST类型API
func ExampleApiPost() {
	sdk := getWxappSdk()

	res, err := sdk.Post("/wxa/query_urllink", nil, map[string]interface{}{
		"url_link": "https://wxaurl.cn/BQZRrcFCPvg",
	})

	fmt.Println(res, err)

	// Output: xxx
}

func getWxappSdk() *wxapp.Wxapp {
	opts := wxapp.Options{
		Debug:  false,
		Appid:  "wx25da2eca8fa3f4ef",
		Secret: "6cca21a5d4b694574fa912a0a928a742",
	}

	return wxapp.New(opts)
}
