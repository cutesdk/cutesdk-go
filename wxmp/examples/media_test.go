package examples

import (
	"fmt"
	"log"
)

func ExampleGetMediaList() {
	client := getClient()

	uri := "/cgi-bin/material/batchget_material"
	params := map[string]interface{}{
		"type":   "news",
		"offset": 0,
		"count":  20,
	}

	res, err := client.PostWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("get media list failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleUploadTempMedia() {
	client := getClient()

	mediaType := "image"
	filePath := "./qrcode.jpg"

	res, err := client.UploadTempMedia(mediaType, filePath)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("upload temp media failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleUploadMedia() {
	client := getClient()

	mediaType := "video"
	filePath := "./video.mp4"
	extra := map[string]interface{}{
		"description": `{"title":"video title", "introduction":"video intro"}`,
	}
	res, err := client.UploadMedia(mediaType, filePath, extra)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("upload image failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleUploadImage() {
	client := getClient()

	filePath := "./qrcode.jpg"
	res, err := client.UploadImage(filePath)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("upload image failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
