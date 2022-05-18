package tests

import (
	"time"

	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/wxpay/v2"
)

var (
	mchId  string
	appid  string
	apiKey string
)

func getClient() *wxpay.Client {
	opts := &wxpay.Options{
		MchId:  mchId,
		Appid:  appid,
		ApiKey: apiKey,
		Request: &request.Options{
			Debug:   true,
			Timeout: 5 * time.Second,
		},
	}

	client, err := wxpay.NewClient(opts)

	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	mchId = "1498014222"
	appid = "wx25da2eca8fa3f4ef"
	apiKey = "Q5xQzBMvdvKQgn3Li3e26XVb4TNYuS13"
}
