package tests

import (
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/ttapp"
)

var (
	appid  string
	secret string
)

func getClient() *ttapp.Client {
	opts := &ttapp.Options{
		Appid:  appid,
		Secret: secret,
		Request: &request.Options{
			Debug: true,
		},
	}

	client, err := ttapp.NewClient(opts)
	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	appid = "ttfb118dac643b1233"
	secret = "92811c680709b5bbc442dba42bb2a681cf04acd5"
}
