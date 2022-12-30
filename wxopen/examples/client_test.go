package examples

import (
	"fmt"
	"log"

	"github.com/cutesdk/cutesdk-go/wxopen"
)

var (
	appid  string
	secret string
)

func ExampleNewClient() {
	fmt.Printf("%T", getClient())

	// Output: *wxopen.Client
}

func getClient() *wxopen.Client {
	cli, _ := wxopen.NewClient(&wxopen.Options{
		Debug:  true,
		Appid:  appid,
		Secret: secret,
	})

	return cli
}

func getWxappClient() *wxopen.WxappClient {
	cli := getClient()

	appid := "xxx"
	refreshToken := "xxx"
	wxappCli, err := cli.NewWxappClient(appid, refreshToken)
	if err != nil {
		log.Fatalf("new wxapp client failed: %v\n", err)
	}

	return wxappCli
}

func getWxmpClient() *wxopen.WxmpClient {
	cli := getClient()

	appid := "xxx"
	refreshToken := "xxx"
	wxmpCli, err := cli.NewWxmpClient(appid, refreshToken)
	if err != nil {
		log.Fatalf("new wxapp client failed: %v\n", err)
	}

	return wxmpCli
}
