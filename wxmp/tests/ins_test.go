package tests

import (
	"log"

	"github.com/cutesdk/cutesdk-go/wxmp"
)

func getIns() *wxmp.Instance {
	opts := &wxmp.Options{
		Appid:  "xxx",
		Secret: "xxx",
		Debug:  true,
	}

	ins, err := wxmp.New(opts)

	if err != nil {
		log.Fatalf("new wxmp sdk failed: %v", err)
	}

	return ins
}
