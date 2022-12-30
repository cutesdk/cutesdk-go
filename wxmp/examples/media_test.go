package examples

import (
	"fmt"
	"log"
)

func ExampleGetMediaList() {
	cli := getClient()

	uri := "/cgi-bin/material/batchget_material"
	params := map[string]interface{}{
		"type":   "news",
		"offset": 0,
		"count":  20,
	}

	res, err := cli.PostWithToken(uri, params)
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
	cli := getClient()

	mediaType := "image"
	filePath := "./qrcode.jpg"

	res, err := cli.UploadTempMedia(mediaType, filePath)
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
	cli := getClient()

	mediaType := "video"
	filePath := "./video.mp4"
	extra := map[string]interface{}{
		"description": `{"title":"video title", "introduction":"video intro"}`,
	}
	res, err := cli.UploadMedia(mediaType, filePath, extra)
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
	cli := getClient()

	filePath := "./qrcode.jpg"
	res, err := cli.UploadImage(filePath)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("upload image failed: %s\n", res)
	}

	fmt.Println("ok")
	// Output: ok
}
