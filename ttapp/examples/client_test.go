package examples

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/ttapp"
)

var (
	appid  string
	secret string
)

func ExampleNewClient() {
	fmt.Printf("%T", getClient())

	// Output: *ttapp.Client
}

func getClient() *ttapp.Client {
	cli, _ := ttapp.NewClient(&ttapp.Options{
		Debug:  true,
		Appid:  appid,
		Secret: secret,
	})

	return cli
}
