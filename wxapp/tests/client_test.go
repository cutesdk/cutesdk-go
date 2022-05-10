package tests

import (
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/wxapp"
)

var (
	appid  string
	secret string
)

func getClient() *wxapp.Client {
	opts := &wxapp.Options{
		Appid:  appid,
		Secret: secret,
		Request: &request.Options{
			Debug: true,
		},
	}
	client, err := wxapp.NewClient(opts)
	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	appid = "wx25da2eca8fa3f4ef"
	secret = "1324f564b26f9f8006515e13660876ef"
}
