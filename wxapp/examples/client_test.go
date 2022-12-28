package examples

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/wxapp"
)

var (
	appid  string
	secret string
)

func ExampleNewClient() {
	fmt.Printf("%T", getClient())

	// Output: *wxapp.Client
}

func getClient() *wxapp.Client {
	client, _ := wxapp.NewClient(&wxapp.Options{
		Debug:  true,
		Appid:  appid,
		Secret: secret,
	})

	return client
}
