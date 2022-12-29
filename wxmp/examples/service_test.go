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

func ExampleSendVideoMsg() {
	client := getClient()

	uri := "/cgi-bin/message/custom/send"
	params := map[string]interface{}{
		"touser":  "xxx",
		"msgtype": "video",
		"video": map[string]interface{}{
			"media_id":       "xxx",
			"thumb_media_id": "xxx",
			"title":          "video title",
			"description":    "video intro",
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

func ExampleSendNewsMsg() {
	client := getClient()

	uri := "/cgi-bin/message/custom/send"
	params := map[string]interface{}{
		"touser":  "xxx",
		"msgtype": "news",
		"news": map[string]interface{}{
			"articles": []map[string]interface{}{
				{
					"title":       "news title",
					"description": "news description",
					"url":         "https://baidu.com",
					"picurl":      "xxx",
				},
			},
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

func ExampleSendMenuMsg() {
	client := getClient()

	uri := "/cgi-bin/message/custom/send"
	params := map[string]interface{}{
		"touser":  "xxx",
		"msgtype": "msgmenu",
		"msgmenu": map[string]interface{}{
			"head_content": "head content",
			"list": []map[string]interface{}{
				{
					"id":      "01",
					"content": "first option",
				},
				{
					"id":      "02",
					"content": "second option",
				},
			},
			"tail_content": "tail content",
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
			"appid":          "xxx",
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
