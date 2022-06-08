package tests

import (
	"github.com/cutesdk/cutesdk-go/wxmp"
)

func getIns() *wxmp.Instance {
	opts := &wxmp.Options{
		Appid:  "wx4833a74fc9337238",
		Secret: "81688583fb7064777a363c4c30eae6d7",
		Debug:  true,
	}

	ins, _ := wxmp.New(opts)

	return ins
}
