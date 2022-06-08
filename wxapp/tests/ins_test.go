package tests

import (
	"github.com/cutesdk/cutesdk-go/wxapp"
)

func getIns() *wxapp.Instance {
	opts := &wxapp.Options{
		Appid:  "wx25da2eca8fa3f4ef",
		Secret: "1324f564b26f9f8006515e13660876ef",
		Debug:  true,
	}
	ins, _ := wxapp.New(opts)

	return ins
}
