package examples

import (
	"fmt"

	"github.com/cutesdk/cutesdk-go/wxwork"
)

var (
	corpid         string
	contactSecret  string
	customerSecret string
	serviceSecret  string
	agentId        string
	agentSecret    string
)

func ExampleNewContactClient() {
	fmt.Printf("%T", getContactClient())

	// Output: *wxwork.Client
}

func getContactClient() *wxwork.Client {
	cli, _ := wxwork.NewClient(&wxwork.Options{
		Corpid: corpid,
		Appid:  "contact",
		Secret: contactSecret,
		Debug:  true,
	})

	return cli
}

func getCustomerClient() *wxwork.Client {
	cli, _ := wxwork.NewClient(&wxwork.Options{
		Corpid: corpid,
		Appid:  "customer",
		Secret: customerSecret,
		Debug:  true,
	})

	return cli
}

func getServiceClient() *wxwork.Client {
	cli, _ := wxwork.NewClient(&wxwork.Options{
		Corpid: corpid,
		Appid:  "service",
		Secret: serviceSecret,
		Debug:  true,
	})

	return cli
}

func getAgentClient() *wxwork.Client {
	cli, _ := wxwork.NewClient(&wxwork.Options{
		Corpid: corpid,
		Appid:  agentId,
		Secret: agentSecret,
		Debug:  true,
	})

	return cli
}
