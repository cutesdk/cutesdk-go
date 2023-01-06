package examples

import (
	"fmt"
	"log"
)

func ExampleUploadTempMedia() {
	cli := getServiceClient()

	mediaType := "image"
	filePath := "./qrcode.jpg"

	res, err := cli.UploadTempMedia(mediaType, filePath)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("upload temp media failed: %s\n", res)
	}

	fmt.Println(res.GetString("media_id"))
	// Output: ok
}

func ExampleGetTempMedia() {
	cli := getServiceClient()

	mediaId := "xxx"

	uri := "/cgi-bin/media/get"
	params := map[string]interface{}{
		"media_id": mediaId,
	}

	res, err := cli.GetWithToken(uri, params)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") > 0 {
		log.Fatalf("create qrcode failed: %s\n", res)
	}

	if err := res.SaveAsFile("./video.mp4"); err != nil {
		log.Fatalf("save qrcode failed: %v\n", err)
	}

	fmt.Println("ok")
	// Output: ok
}

func ExampleUploadImage() {
	cli := getServiceClient()

	filePath := "./qrcode.jpg"
	res, err := cli.UploadImage(filePath)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("upload image failed: %s\n", res)
	}

	fmt.Println(res.GetString("url"))
	// Output: ok
}

func ExampleUploadVideo() {
	cli := getServiceClient()

	mediaType := "video"
	filePath := "./video.mp4"

	res, err := cli.UploadTempMedia(mediaType, filePath)
	if err != nil {
		log.Fatalf("request api failed: %v\n", err)
	}

	if res.GetInt("errcode") != 0 {
		log.Fatalf("upload image failed: %s\n", res)
	}

	fmt.Println(res.GetString("media_id"))
	// Output: ok
}
