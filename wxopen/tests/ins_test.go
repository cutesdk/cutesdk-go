package tests

import (
	"github.com/cutesdk/cutesdk-go/wxopen"
)

var (
	componentAppid     string
	componentAppsecret string
	verifyToken        string
	encodingAesKey     string
)

func getIns() *wxopen.Instance {
	opts := &wxopen.Options{
		ComponentAppid:     componentAppid,
		ComponentAppsecret: componentAppsecret,
		VerifyToken:        verifyToken,
		EncodingAesKey:     encodingAesKey,
		Debug:              true,
	}

	ins, err := wxopen.New(opts)

	if err != nil {
		panic(err)
	}

	return ins
}

func init() {
	componentAppid = "xxx"
	componentAppsecret = "xxx"
	verifyToken = "xxx"
	encodingAesKey = "xxx"
}
