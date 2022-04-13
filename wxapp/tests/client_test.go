package tests

import (
	"github.com/cutesdk/cutesdk-go/common/app"
	"github.com/cutesdk/cutesdk-go/wxapp"
)

var (
	appid  string
	secret string
)

func getClient() *wxapp.Client {
	client, err := wxapp.NewClient(appid, secret, app.WithDebug(true))
	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	appid = "wx25da2eca8fa3f4ef"
	secret = "1324f564b26f9f8006515e13660876ef"
}
