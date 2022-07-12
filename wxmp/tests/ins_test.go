package tests

import (
	"log"

	"github.com/cutesdk/cutesdk-go/wxmp"
)

func getIns() *wxmp.Instance {
	opts := &wxmp.Options{
		Appid:  "wx8dcd98079e13d33f",
		Secret: "76f2af7b3c53826e88b0dad8eb2e3e77",
		Debug:  true,
	}

	ins, err := wxmp.New(opts)

	if err != nil {
		log.Fatalf("new wxmp sdk failed: %v", err)
	}

	return ins
}
