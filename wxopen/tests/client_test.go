package tests

import (
	"time"

	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/common/request"
	"github.com/cutesdk/cutesdk-go/wxopen"
)

var (
	componentAppid     string
	componentAppsecret string
)

func getClient() *wxopen.Client {
	opts := &wxopen.Options{
		ComponentAppid:     componentAppid,
		ComponentAppsecret: componentAppsecret,
		Request: &request.Options{
			Debug:   true,
			Timeout: 5 * time.Second,
		},
		Cache: &cache.Options{
			Driver: "redis",
			Conf: map[string]interface{}{
				"dsn":     "redis://:@127.0.0.1:6379/1",
				"timeout": "3s",
			},
		},
	}

	client, err := wxopen.NewClient(opts)

	if err != nil {
		panic(err)
	}

	return client
}

func init() {
	componentAppid = "wxf2f955ce09390e6a"
	componentAppsecret = "d6e9032e5f5bcea2f96b66f2c4e1cab8"
}
