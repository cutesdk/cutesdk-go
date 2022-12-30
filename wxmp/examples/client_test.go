package examples

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/wxmp"
)

var (
	appid  string
	secret string
)

func ExampleNewClient() {
	fmt.Printf("%T", getClient())

	// Output: *wxmp.Client
}

func getClient() *wxmp.Client {
	cli, _ := wxmp.NewClient(&wxmp.Options{
		Debug:  true,
		Appid:  appid,
		Secret: secret,
	})

	return cli
}
