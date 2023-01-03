package examples

import (
	"fmt"
	"log"

	"github.com/cutesdk/cutesdk-go/wxpay/v3"
)

var (
	mchId    string
	apiKey   string
	serialNo string
	keyPem   string
	keyPath  string
	appid    string
)

func ExampleNewClient() {
	fmt.Printf("%T", getClient())

	// Output: *wxpay.Client
}

func getClient() *wxpay.Client {
	cli, err := wxpay.NewClient(&wxpay.Options{
		MchId:    mchId,
		ApiKey:   apiKey,
		SerialNo: serialNo,
		KeyPem:   keyPem,
		KeyPath:  keyPath,
		Debug:    true,
	})

	if err != nil {
		log.Fatalf("new wxpay client failed: %v\n", err)
	}

	return cli
}
