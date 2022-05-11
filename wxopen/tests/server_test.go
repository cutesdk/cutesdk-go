package tests

import (
	"github.com/cutesdk/cutesdk-go/common/cache"
	"github.com/cutesdk/cutesdk-go/wxopen"
)

func getServer() *wxopen.Server {
	opts := &wxopen.Options{
		ComponentAppid:     componentAppid,
		ComponentAppsecret: componentAppsecret,
		VerifyToken:        verifyToken,
		EncodingAesKey:     encodingAesKey,
		Cache: &cache.Options{
			Driver: "redis",
			Conf: map[string]interface{}{
				"dsn":     "redis://:@127.0.0.1:6379/1",
				"timeout": "3s",
			},
		},
	}

	server, err := wxopen.NewServer(opts)

	if err != nil {
		panic(err)
	}

	return server
}
