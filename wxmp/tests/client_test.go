package tests

import (
	"github.com/cutesdk/cutesdk-go/common/app"
	"github.com/cutesdk/cutesdk-go/wxmp"
)

var (
	appid  string
	secret string
)

func getClient() *wxmp.Client {
	client, err := wxmp.NewClient(appid, secret, app.WithDebug(true))
	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	appid = "wx4833a74fc9337238"
	secret = "81688583fb7064777a363c4c30eae6d7"
}
