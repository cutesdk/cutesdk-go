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

	ins, err := ttapp.New(opts)
	if err != nil {
		panic(err)
	}

	return ins
}

func init() {
	appid = "xxx"
	secret = "xxx"
}
