package tests

import (
	"github.com/cutesdk/cutesdk-go/wxapp"
)

func getIns() *wxapp.Instance {
	opts := &wxapp.Options{
		Appid:  "wx25da2eca8fa3f4ef",
		Secret: "371bcf903656f33218b671f325bc6d1e",
		Debug:  true,
	}
	ins, _ := wxapp.New(opts)

	return ins
}
