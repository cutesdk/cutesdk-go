package examples

import (
	"fmt"
	"log"

	"github.com/cutesdk/cutesdk-go/wxpay/v2"
)

var (
	mchId    string
	apiKey   string
	certPath string
	keyPath  string
	certPem  string
	keyPem   string
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
		CertPath: certPath,
		KeyPath:  keyPath,
		CertPem:  certPem,
		KeyPem:   keyPem,
		Debug:    true,
	})

	if err != nil {
		log.Fatalf("new wxpay client failed: %v\n", err)
	}

	return cli
}
