package tests

import (
	"github.com/cutesdk/cutesdk-go/common/app"
	"github.com/cutesdk/cutesdk-go/ttapp"
)

var (
	appid  string
	secret string
)

func getClient() *ttapp.Client {
	client, err := ttapp.NewClient(appid, secret, app.WithDebug(true))
	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	appid = "ttfb118dac643b1233"
	secret = "92811c680709b5bbc442dba42bb2a681cf04acd5"
}
