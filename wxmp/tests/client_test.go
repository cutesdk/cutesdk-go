package tests

import (
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/wxmp"
)

var (
	appid  string
	secret string
)

func getClient() *wxmp.Client {
	opts := &wxmp.Options{
		Appid:  appid,
		Secret: secret,
		Request: &request.Options{
			Debug: true,
		},
	}
	client, err := wxmp.NewClient(opts)
	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	appid = "wx4833a74fc9337238"
	secret = "81688583fb7064777a363c4c30eae6d7"
}
