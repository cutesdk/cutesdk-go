package tests

import (
	"github.com/cutesdk/cutesdk-go/wxapp"
)

func getIns() *wxapp.Instance {
	opts := &wxapp.Options{
		Appid:  "xxx",
		Secret: "xxx",
		Debug:  true,
	}
	ins, _ := wxapp.New(opts)

	return ins
}
