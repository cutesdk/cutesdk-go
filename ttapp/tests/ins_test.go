package tests

import (
	"github.com/cutesdk/cutesdk-go/ttapp"
)

var (
	appid  string
	secret string
)

func getIns() *ttapp.Instance {
	opts := &ttapp.Options{
		Appid:  appid,
		Secret: secret,
		Debug:  true,
	}

	ins, _ := ttapp.New(opts)

	return ins
}

func init() {
	appid = "ttfb118dac643b1233"
	secret = "92811c680709b5bbc442dba42bb2a681cf04acd5"
}
