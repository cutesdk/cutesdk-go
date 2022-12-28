package examples

import (
	"fmt"
	"log"
)

func ExampleSendTextMsg() {
	client := getClient()

	uri := "/cgi-bin/message/custom/send"
	params := map[string]interface{}{
		"touser":  "xxx",
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": "hello",
		},
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("send service text msg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleSendImageMsg() {
	client := getClient()

	uri := "/cgi-bin/message/custom/send"
	params := map[string]interface{}{
		"touser":  "xxx",
		"msgtype": "image",
		"image": map[string]interface{}{
			"media_id": "xxx",
		},
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("send service text msg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleSendLinkMsg() {
	client := getClient()

	uri := "/cgi-bin/message/custom/send"
	params := map[string]interface{}{
		"touser":  "xxx",
		"msgtype": "link",
		"link": map[string]interface{}{
			"title":       "link title",
			"description": "link description",
			"url":         "https://baidu.com",
			"thumb_url":   "xxx",
		},
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("send service text msg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleSendMiniprogramMsg() {
	client := getClient()

	uri := "/cgi-bin/message/custom/send"
	params := map[string]interface{}{
		"touser":  "xxx",
		"msgtype": "miniprogrampage",
		"miniprogrampage": map[string]interface{}{
			"title":          "miniprogram title",
			"pagepath":       "pages/index/index",
			"thumb_media_id": "xxx",
		},
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("send service text msg failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
