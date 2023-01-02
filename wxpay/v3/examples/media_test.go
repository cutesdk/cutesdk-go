package examples

import (
	"fmt"
	"log"
)

func ExampleUploadImage() {
	cli := getClient()

	filePath := "./qrcode.jpg"
	res, err := cli.UploadImage(filePath)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %d, %s, %s\n", wxerr.StatusCode, wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("media_id"))
	// Output: xxx
}

func ExampleUploadMarketingImage() {
	cli := getClient()

	filePath := "./qrcode.jpg"
	res, err := cli.UploadMarketingImage(filePath)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %d, %s, %s\n", wxerr.StatusCode, wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("media_url"))
	// Output: xxx
}

func ExampleUploadMerchantImage() {
	cli := getClient()

	filePath := "./qrcode.jpg"
	res, err := cli.UploadMerchantImage(filePath)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %d, %s, %s\n", wxerr.StatusCode, wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("media_id"))
	// Output: xxx
}

func ExampleUploadVideo() {
	cli := getClient()

	filePath := "./video.mp4"
	res, err := cli.UploadVideo(filePath)
	if err != nil {
		if wxerr := cli.UnwrapError(err); wxerr != nil {
			log.Fatalf("request api failed: %d, %s, %s\n", wxerr.StatusCode, wxerr.Code, wxerr.Message)
		}
		log.Fatalf("request api failed: %v\n", err)
	}

	fmt.Println(res.GetString("media_id"))
	// Output: xxx
}
