package examples

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/qqapp"
)

var (
	appid  string
	secret string
)

func ExampleNewClient() {
	fmt.Printf("%T", getClient())

	// Output: *qqapp.Client
}

func getClient() *qqapp.Client {
	cli, _ := qqapp.NewClient(&qqapp.Options{
		Debug:  true,
		Appid:  appid,
		Secret: secret,
	})

	return cli
}
