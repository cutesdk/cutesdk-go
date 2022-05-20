package tests

import (
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/wxwork"
)

var (
	corpid        string
	serviceSecret string
)

func getServiceClient() *wxwork.Client {
	opts := &wxwork.Options{
		Corpid:  corpid,
		Agentid: "service",
		Secret:  serviceSecret,
		Request: &request.Options{
			Debug: true,
		},
	}
	client, err := wxwork.NewClient(opts)
	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	corpid = "wwa3f1494ad3d3713d"
	serviceSecret = "44d2imiTA4EhySj2TVfstu6LRh4dyGZef8oQcb43n_Y"
}
